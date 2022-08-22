package hello

import (
	"fmt"
)

type myLexer struct {
	//待解析指令暂时存input里面，用byte数组方便逐字符操作
	input []byte
	//当前扫描到字符的索引
	index int
	//存储已经扫描到的内容，当内容为hello时，返回HELLO
	buffer string
}

func (m *myLexer) Lex(lval *helloSymType) int {
	if len(m.input) == 0 {
		return 0 //表示已经解析完所有的内容了
	}
	for i := 0; i < len(m.input); i++ {
		char := m.input[i]
		switch char {
		case 'h', 'e', 'l', 'o':
			m.buffer += string(char)
		}
		if m.buffer == "hello" {
			return HELLO
		}
	}
	return 0
}

func (m *myLexer) Error(s string) {
	fmt.Println(s)
}
