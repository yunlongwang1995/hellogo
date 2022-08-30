package __thread

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSync1(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("子线程...")
		wg.Done()
	}()
	fmt.Println("主线程")
	wg.Wait()
}
