package main

import "fmt"

func main() {
	n1 := 100
	fmt.Printf(" n1的类型:%T\n n1的值:%v\n n1的地址:%v\n", n1, n1, &n1)
	fmt.Println()
	// n1的类型:int
	// n1的值:100
	// n1的地址:0xc00001c0d0
	n2 := new(int) ////new返回的是指针类型
	fmt.Printf(" n2的类型:%T\n n2的值:%v\n n2的地址:%v\n n2指向的值:%v\n", n2, n2, &n2, *n2)

	// n2的类型:*int
	// n2的值:0xc00001c0d8
	// n2的地址:0xc000012030
	// n2指向的值:0
	*n2 = 12
	fmt.Println(*n2)
}
