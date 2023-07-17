package main

import "fmt"

// 支持1到多个参数
// 可变参数必须放在形参列表的最后，即： args ...int 放在后面
func sum(a int, args ...int) int {
	s := a
	for i := 0; i < len(args); i++ {
		s += args[i]
	}
	return s
}

// 支持0到多个参数
func sum2(args ...int) int {
	s := 0
	for i := 0; i < len(args); i++ {
		s += args[i]
	}
	return s
}
func main() {
	res := sum(10, 44, 12, 2)
	res2 := sum2()
	res3 := sum2(10, 44, 12, 2)
	fmt.Println(res)
	fmt.Println(res2)
	fmt.Println(res3)
}

//args是silce切片（动态数组）
