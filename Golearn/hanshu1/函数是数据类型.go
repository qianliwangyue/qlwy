package main

import "fmt"

func ssum(a, b int) int {
	return a + b
}

func main() {
	a := ssum
	fmt.Printf("a的类型:%T\nSSum的类型:%T\n", a, ssum)
	//	a的类型:func(int, int) int
	//	SSum的类型:func(int, int) int
	ss := a(15, 5)
	fmt.Println(ss) //20

	res := aa(ssum, 51, 21)
	fmt.Println(res) //72
}

// 函数本身是一种数据类型，因此可以作为另一个函数的形参。
// func(int, int) int ==>代表数据类型
func aa(Ass func(int, int) int, num1 int, num2 int) int {
	return Ass(num1, num2)
}
