package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	a := cap(ch) //获取make函数的容积
	fmt.Println("hhhh", a)
	for i := 1; i < 6; i++ {
		ch <- i //将1~5逐个送入通道
	}
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch) //5次打印，接收通道里的内容：1~5
	}
}

//送入通道的值的数量不能超出（容积）即5
//接收的值的数量不能超过发送的数量
