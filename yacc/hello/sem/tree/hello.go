package tree

import (
	"fmt"
	"time"
)

type HelloStatement struct {
	Name string
}

func (hello *HelloStatement) Exec() {
	fmt.Println("time: " + time.Now().String())
	fmt.Println("hello: " + hello.Name)
}
