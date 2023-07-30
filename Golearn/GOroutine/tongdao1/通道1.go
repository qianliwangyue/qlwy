// 不带缓存的通道必须发送和接收同步进行
package main

import "fmt"

func main() {
	msg := make(chan string) //创建string类型的通道msg

	type AA struct {
		A string
		B int
	}
	msg1 := make(chan *AA) //创建结构体类型的通道msg1
	var m AA
	m.A = "不敢相信"
	m.B = 666

	go func() {
		msg <- "哈哈哈" //将"哈哈哈"传入通道msg
		msg1 <- &m
	}()
	m1 := <-msg //定义变量m1接收通道msg传来的内容
	m2 := <-msg1
	fmt.Println(m1, m2, *m2)
}
