package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once //定义一个sync.Once对象
	onceBody := func() {
		fmt.Println("test only once,这里只打印一次！")
	} //定义一个闭包函数赋给onceBody
	done := make(chan bool) //定义一个布尔类型的通道done
	for i := 0; i < 6; i++ {
		go func() {
			once.Do(onceBody) //once.Do只执行一次（用来初始化）
			done <- true
		}()
	}
	for i := 0; i < 6; i++ {
		//<-done
		fmt.Println(<-done)
	}
}
