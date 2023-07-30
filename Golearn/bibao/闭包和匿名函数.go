package main

import (
	"fmt"
	"strings"
)

// 函数名/参数a       /返回值：函数
func makeS(a string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, a) { //HasSuffix函数（参数1：文件名，参数2：后缀名）判断参数1的后缀名，是否==参数2，返回bool
			return name + a
		}
		return name
	}
}
func aadd() func(int) int { //返回值为函数func(int) int，和外面的变量n共同构成闭包
	var n = 10 //闭包是类，函数是操作，n是字段，n初始化一次
	return func(x int) int {
		n += x
		return n
	}
}
func main() {
	f2 := makeS(".jpg")
	fmt.Println("文件名处理后=", f2("aaa"))
	fmt.Println("文件名处理后=", f2("bbb.jpg"))
	//文件名处理后= aaa.jpg
	//文件名处理后= bbb.jpg

	v := aadd()
	fmt.Println(v(1)) //11
	fmt.Println(v(2)) //13
	fmt.Println(v(3)) //16
}
