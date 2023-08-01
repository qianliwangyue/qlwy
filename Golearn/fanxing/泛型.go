package main

import "fmt"

// 泛型表示广泛的类型（非单一类型）即使函数可以支持多种形参类型
// 方括号声明参数类型：int或float64
// 方括号中any表示任意类型
// T可用其他字母代替
func getMax[T int | float64](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}

}

// 定义一个接口，存储泛型包含的类型
type shu interface {
	//uint8、uint16、uint32、float32、int64及其衍生类型（myint）
	uint8 | uint16 | uint32 | float32 | ~int64
}
type myint int64 //定义一个新类型（int64衍生类型）

var a, b myint

func getMax1[T shu](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}

}
func main() {

	a = 1649849849849849849
	b = 2165165165165156151

	//fmt.Println(getMax[int](4, 5))
	//fmt.Println(getMax[float64](4.52, 5.36))
	//调用泛型函数时，可省略数据类型
	fmt.Println(getMax(4, 5))
	fmt.Println(getMax(4.12, 8.13))
	fmt.Println(getMax1(a, b))
}
