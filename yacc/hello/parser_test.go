package hello

import (
	"hellogo/yacc/hello/sem/tree"
	"testing"
	"time"
)

func TestParser(t *testing.T) {
	//正确使用
	helloParse(&myLexer{
		input:  []byte("hello"),
		index:  0,
		buffer: "",
	})

	time.Sleep(time.Second)

	//未定义语法，会报错
	helloParse(&myLexer{
		input:  []byte("hhhhhh"),
		index:  0,
		buffer: "",
	})
}

func TestHello(t *testing.T) {
	st := tree.HelloStatement{
		Name: "parser...",
	}

	st.Exec()
}
