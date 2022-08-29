package main

/*
#include <stdio.h>

void printint(int v) {
    printf("printint: %d\n", v);
}
*/
import "C"
import "fmt"

func main() {
	// CGO 测试
	v := 42
	C.printint(C.int(v))
	fmt.Println("hello cgo example")

	// 8进制与16进制
	const hex = '\x11'
	fmt.Println(hex)
	const octal = '\011'
	fmt.Println(octal)
}
