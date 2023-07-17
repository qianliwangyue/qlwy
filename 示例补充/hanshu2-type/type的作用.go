package main

import "fmt"

// 定义一个函数
func addd(a int, b int) int {
	return a + b
}
func co(a, b int) int {
	return (a - b)
}

// 将函数定义为一个名为add的类型
type add func(int, int) int

// 将add类型作为形参赋值给函数
func Hff(as add, n1 int, n2 int) int {
	return as(n1, n2)
}
func main() {
	re1 := Hff(addd, 51, 6) //57
	re2 := Hff(co, 55, 41)  //14
	fmt.Println(re1, re2)

}
