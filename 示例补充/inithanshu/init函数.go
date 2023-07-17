package inithanshu

import "fmt"

func init() {
	fmt.Println("欢迎光临！--1")
}
func init() {
	fmt.Println("欢迎光临！--2")
}
func Add(a int, b int) int {
	return a + b
}

//init 函数没有参数和返回值，用于初始化
//一个包内可以包含多个init函数
