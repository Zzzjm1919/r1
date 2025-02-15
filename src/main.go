package main

/*
#include <stdio.h>

void print_hello() {
    printf("Hello, World from C!\n");
}
*/
import "C"
import "fmt"

func main() {
	C.print_hello() // 调用 C 函数
	fmt.Println("======")
}
