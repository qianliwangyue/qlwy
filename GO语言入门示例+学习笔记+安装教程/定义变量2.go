package main

import "fmt"

var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

// 这种不带声明格式的只能在函数体中出现
// g, h := 123, "hello"
func main() {
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)
	var a string = "abc"
	fmt.Println("hello, world", a)

	fmt.Println("========================================")

	var j, n int
	var m string
	j, n, m = 5, 7, "abc"
	//等价于//	j, n, m := 5, 7, "abc"
	fmt.Println(j, n, m)

	fmt.Println("========================================")
	var u, i uint8 = 5, 6
	u, i = i, u //将u和i值互换
	fmt.Println(u, i)
	_, i = 3, 8 //下划线表示对应值被舍弃
	fmt.Println(i)

	fmt.Println("========================================")
	//下划线表示对应值被舍弃，实例：
	_, numb, strs := numbers() //只获取函数返回值的后两个
	fmt.Println(numb, strs)
}

//主函数外，定义函数number
func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	b++
	return a, b, c
}
