package main

import "fmt"

//匿名函数getSequence作为函数名, func()表示参数列表，int为返回值类型
func gf() func() int {
	//1
	i := 0
	return func() int {
		i += 1
		return i
	}
	//2
	//1~2之间是一个闭包
}

func main() {
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := gf()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	c := gf() /* c 为一个函数，函数 i 为 0 */
	fmt.Println(c())
	fmt.Println(c())
}
