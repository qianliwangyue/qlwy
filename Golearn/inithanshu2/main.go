package main

import (
	"fmt"

	"Golearn/inithanshu"
)

func main() {
	c := inithanshu.Add(5, 6)
	fmt.Println(c)
}

/*
运行结果：
欢迎光临！
11

GOLearn/inithanshu/init函数.go
中
func init() {
	fmt.Println("欢迎光临！")
}
init 函数没有参数和返回值，用于初始化，同一个包内可以包含多个init函数


如果只需要GOLearn/inithanshu/init函数.go中的init用来初始化，
而不需要使用init函数.go中的函数，则可以在导入包时，前面加上下划线

*/
