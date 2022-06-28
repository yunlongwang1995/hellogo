package __io

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"testing"
	"unsafe"
)

type zn_iocb struct {
	/* these are internal to the kernel/libc. */
	aio_data      uint64
	aio_key       uint32
	aio_reserved1 uint32
	/* common fields */
	aio_lio_opcode uint16 /* see IOCB_CMD_ above */
	aio_reqprio    int16
	aio_fildes     uint32
	aio_buf        uint64
	aio_nbytes     uint64
	aio_offset     int64
	/* extra parameters */
	aio_reserved2 uint64
	aio_flags     uint32
	aio_resfd     uint32
} /* 64 bytes */

type zn_io_event struct {
	data uint64 /* the data field from the iocb */
	obj  uint64 /* what iocb this event came from */
	res  int64  /* result code for this event */
	res2 int64  /* secondary result */
}

type zn_timespec struct {
	tv_sec  uint64 /* seconds */
	tv_nsec uint64 /* nanoseconds */
}

const (
	AlignSize              = 4096
	IOCB_CMD_PREAD  uint16 = 0
	IOCB_CMD_PWRITE uint16 = 1
	IOCB_CMD_FSYNC  uint16 = 2
	IOCB_CMD_FDSYNC uint16 = 3
	// Minimum block size
	BlockSize = 4096

	NR_io_setup     = 206
	NR_io_destroy   = 207
	NR_io_getevents = 208
	NR_io_submit    = 209
	NR_io_cancel    = 210
)

func alignment(block []byte, AlignSize int) int {
	return int(uintptr(unsafe.Pointer(&block[0])) & uintptr(AlignSize-1))
}

// AlignedBlock returns []byte of size BlockSize aligned to a multiple of AlignSize in memory (must be power of two)
func AlignedBlock(BlockSize int) []byte {
	block := make([]byte, BlockSize+AlignSize)
	if AlignSize == 0 {
		return block
	}
	a := alignment(block, AlignSize)
	offset := 0
	if a != 0 {
		offset = AlignSize - a
	}
	block = block[offset : offset+BlockSize]
	// Can't check alignment of a zero sized block
	if BlockSize != 0 {
		a = alignment(block, AlignSize)
		if a != 0 {
			log.Fatal("Failed to align block")
		}
	}
	return block
}

func testReadWrite(fp *os.File) {
	block1 := AlignedBlock(BlockSize)
	for i := 0; i < len(block1); i++ {
		block1[i] = 66
	}

	_, err := fp.Write(block1)
	if err != nil {
		fmt.Println("Failed to write", err)
		return
	}

	// Read the file
	block2 := AlignedBlock(BlockSize)
	_, err = fp.ReadAt(block2, 0)
	if err != nil {
		fmt.Println("Failed to read", err)
		return
	}

	for i := 0; i < len(block2); i++ {
		fmt.Print(block2[i])
	}
}

func TestAIO(t *testing.T) {
	fp, err := os.OpenFile("./aio.txt", os.O_CREATE|os.O_RDWR|syscall.O_DIRECT, 0666)
	if err != nil {
		fmt.Println("Open err =", err)
		return
	}
	fmt.Println("Open file success", err)
	//testReadWrite(fp)

	// testReadWrite(fp)
	var ctx uint64
	ret, _, err := syscall.Syscall(NR_io_setup, 1024, uintptr(unsafe.Pointer(&ctx)), 0)
	fmt.Println("ret is :", ret, err, ", ctx is: ", ctx)
	if ret < 0 {
		fmt.Println("Failed to NR_io_setup", err)
		return
	}

	buf := AlignedBlock(AlignSize)
	for i := 0; i < len(buf); i++ {
		buf[i] = '0'
	}
	bufptr := &buf[0]
	cur_iocb := new(zn_iocb)
	cur_iocb.aio_buf = *(*uint64)(unsafe.Pointer(&bufptr))
	cur_iocb.aio_reqprio = 0
	cur_iocb.aio_fildes = uint32(fp.Fd())
	cur_iocb.aio_lio_opcode = IOCB_CMD_PWRITE
	cur_iocb.aio_offset = AlignSize * 3
	cur_iocb.aio_nbytes = AlignSize

	cur_iocb.aio_data = 0
	cur_iocb.aio_key = 0
	cur_iocb.aio_reserved1 = 0
	cur_iocb.aio_reserved2 = 0
	cur_iocb.aio_flags = 0
	cur_iocb.aio_resfd = 0

	{
		type SliceMock struct {
			addr uintptr
			len  int
			cap  int
		}
		Len := unsafe.Sizeof(*cur_iocb)
		testBytes := &SliceMock{
			addr: uintptr(unsafe.Pointer(cur_iocb)),
			cap:  int(Len),
			len:  int(Len),
		}
		fmt.Println("iocb memory is:")
		data := *(*[]byte)(unsafe.Pointer(testBytes))
		for i := 0; i < int(Len); i++ {
			fmt.Print(uint8(data[i]), " ")
		}
		fmt.Println()
	}

	ret, _, err = syscall.Syscall(NR_io_submit, uintptr(ctx), 1, uintptr(unsafe.Pointer(&cur_iocb))) // 用不到的就补上 0
	fmt.Println("ret is :", ret, err)
	if err.Error() != "errno 0" {
		fmt.Println("Failed to NR_io_submit: ", err)
		fp.Close()
		ret, _, err = syscall.Syscall(NR_io_destroy, uintptr(ctx), 0, 0)
		return
	}

	var ts zn_timespec
	ts.tv_nsec = 0
	ts.tv_sec = 10
	var eves [10]zn_io_event

	ret, _, err = syscall.Syscall6(NR_io_getevents, uintptr(ctx), 1, 10,
		uintptr(unsafe.Pointer(&eves[0])), uintptr(unsafe.Pointer(&ts)), 0)
	if ret < 0 {
		fmt.Println("Failed to NR_io_getevents", err)
		return
	}
	fmt.Println("ret events: ", uint32(ret), err)

	fp.Close()
	ret, _, err = syscall.Syscall(NR_io_destroy, uintptr(ctx), 0, 0)
	fmt.Println("ret destroy: ", uint32(ret), err)
}
