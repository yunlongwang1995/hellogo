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
	v := 42
	C.printint(C.int(v))

	fmt.Println("hello cgo example")
}
