package main

import "fmt"

func main() {

	p1 := hh1()
	fmt.Println("p1:", p1) //p1: 0xc000014270   显示地址
	fmt.Println("p1", *p1) //p1 深圳
}

// 定义一个返回指针类型的函数
func hh1() *string {
	city := "深圳"
	ptr := &city
	return ptr
}
