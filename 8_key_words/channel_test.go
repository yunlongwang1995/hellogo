package examples

import (
	"fmt"
	"testing"
	"time"
)

func Test_channel_nil(t *testing.T) {
	// case1: channel is nil
	var ch chan int
	go func() {
		for {
			time.Sleep(3 * time.Second)
			fmt.Println("internal...")
		}
	}()
	ch <- 1
	//close(ch)
	fmt.Println("hello end...")
}

func Test_channel_normal(t *testing.T) {

}

func Test_channel_close(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	close(ch)

	ch <- 2

	go func() {
		for {
			time.Sleep(3 * time.Second)
			fmt.Println("internal...")
		}
	}()

	x := <-ch
	fmt.Println(x)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func Test_channel_range(t *testing.T) {
	ch := make(chan int, 1)

	go func() {
		for {
			ch <- 1
			time.Sleep(3 * time.Second)
		}
	}()

	for c := range ch {
		fmt.Println(c)
		fmt.Println("hello...")
	}
}
