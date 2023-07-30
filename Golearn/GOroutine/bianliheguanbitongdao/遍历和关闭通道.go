package main

import (
	"fmt"
)

func fibonacci(n int, ch chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

func main() {
	ch := make(chan int, 6)
	go fibonacci(cap(ch), ch) //cap(ch)：获取make函数定义的容积：6
	for j := range ch {
		fmt.Println(j)
	}
}
