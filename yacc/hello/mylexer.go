package main

import "fmt"

type MyLexer struct {
	input  []byte
	index  int
	buffer string
}

func (m *MyLexer) Lex(lval *helloSymType) int {
	if len(m.input) == 0 {
		return 0
	}

	for _, elem := range m.input {
		switch elem {
		case 'h', 'i', 'e', 'l', 'o':
			m.buffer += string(elem)
		}

		if m.buffer == "hello" {
			return HELLO
		}

		if m.buffer == "hi" {
			return HI
		}
	}

	return 0
}

func (m *MyLexer) Error(s string) {
	fmt.Println("syntax error...", s)
}

func main() {
	// normal case
	helloParse(&MyLexer{input: []byte("hello")})
	helloParse(&MyLexer{input: []byte("hi")})

	// unnormal case
	helloParse(&MyLexer{input: []byte("haha")})
}
