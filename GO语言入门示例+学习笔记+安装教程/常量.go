package main

import (
	"fmt"
	"unsafe"
)

const (
	A = "abc"
	B = len(A)
	C = unsafe.Sizeof(A) //字节数
)

func main() {
	const b string = "abc" //常量b不可变
	fmt.Println(b)

	const LENGTH, WIDTH int = 10, 5
	var area int
	const a, f, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, f, c)
	println(A, B, C)
}
