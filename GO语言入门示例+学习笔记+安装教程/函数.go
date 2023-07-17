package main

import (
	"fmt"
)

//import "fmt"

func main() {
	//var re uint16
	re := max(15, 68)
	fmt.Printf("较大的是：%d\n", re) //自定义格式化输出
	fmt.Println("较大的是：", re)    //换行输出
	fmt.Print("较大的是：", re)      //默认输出，不换行
	fmt.Println()
	fmt.Println("=========================================================")

	a, b, c := "我", "喜欢", "项宇"
	D, E := fd(a, b, c)
	fmt.Println(D, E)

	fmt.Println("=========================================================")

	x, y, z := "我", "喜欢", "项宇"
	x1, y1, z1 := jh(x, y, z) //函数实现值之间的交换
	fmt.Println(x1, y1, z1)

}

// 函数名（参数列表）返回值类型
func max(n1, n2 uint16) uint16 {
	var result uint16
	if n1 > n2 {
		result = n1
	} else {
		result = n2
	}
	return result
}
func fd(a, b, c string) (string, string) {
	return a, c

}
func jh(a, b, c string) (string, string, string) {

	a, c = c, a
	return a, b, c

}
