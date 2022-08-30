package __io

import (
	"fmt"
	"os"
	"testing"
)

// 创建文件
func TestFileCreate(t *testing.T) {
	file, err := os.OpenFile("hellogo.txt", os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("create file failed: ", err)
	}

	if err = file.Close(); err != nil {
		fmt.Println("file close failed: ", err)
	}
}

// 写文件
func TestFileWrite(t *testing.T) {
	file, err := os.OpenFile("hellogo.txt", os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("open failed: ", err)
		return
	}

	_, err = file.Write([]byte("hello go\n"))
	_, err = file.Write([]byte("i am from shandong\n"))
	if err != nil {
		fmt.Println("write failed: ", err)
		return
	}

	if err = file.Close(); err != nil {
		fmt.Println("close failed: ", err)
		return
	}
}

// 读文件
func TestFileRead(t *testing.T) {
	file, err := os.OpenFile("hellogo.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println("open failed: ", err)
		return
	}

	var data = make([]byte, 1024)
	n, err := file.Read(data)
	if err != nil {
		fmt.Println("read failed: ", err)
		return
	}

	fmt.Println(n)
	fmt.Println(string(data[:n]))

	if err = file.Close(); err != nil {
		fmt.Println("close failed: ", err)
		return
	}
}

// 删除文件
func TestFileDelete(t *testing.T) {
	err := os.Remove("hellogo.txt")
	if err != nil {
		fmt.Println("delete failed: ", err)
	}
}
