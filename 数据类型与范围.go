package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num = 28
	var num2 uint8 = 200
	fmt.Printf("num2的类型：%T\n", num2)
	fmt.Println("num字节数：", unsafe.Sizeof(num))

}
