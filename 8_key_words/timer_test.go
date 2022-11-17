package examples

import (
	"fmt"
	"testing"
)
import "time"

func Test_timer(t *testing.T) {
	// now
	fmt.Println(time.Now().String())
	timer := time.NewTimer(5 * time.Second)
	<-timer.C

	// now + 5

	time.AfterFunc(2*time.Second, func() {
		fmt.Println(time.Now().String())
	})
	//fmt.Println("stop timer: ", timer2.Stop())

	timeCh := time.After(5 * time.Second)
	<-timeCh

	// now + 10

	fmt.Println(time.Now().String())
}
