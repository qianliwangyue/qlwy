package main

import (
	"fmt"
)

func sum(a, b int) int {
	defer fmt.Println("a=", a) //a压入栈,a的值10也拷贝压入
	defer fmt.Println("b=", b) //b压入栈，b的值20也拷贝压入
	a++                        //11
	b++                        //21
	res := a + b
	fmt.Println("res1=", res)
	return res
	//执行过后// b出栈// a出栈

}

func main() {
	res := sum(10, 20) //defer,a入栈，b入栈
	//res1=32
	//b= 20
	//a= 10
	fmt.Println("res=", res)
	//res1=32

}
