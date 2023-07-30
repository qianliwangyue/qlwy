package main

import (
	"fmt"
)

func Intcreat() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i //直到通道关闭索要数据才把i添加进通道
		}
	}()
	return ch
}
func main() {
	gen := Intcreat() //不断向通道传入整数0，1，2，3......
	for i := 0; i < 100; i++ {
		fmt.Println(<-gen) //打印接收到的100次的数
	}
}
