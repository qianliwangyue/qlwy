package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	Tongxvn()
	//利用channel的阻塞机制，等待退出信号
	ch:=make(chan os.Signal,0)
	signal.Notify(ch,os.Interrupt,os.Kill)
	<-ch
}

func Tongxvn() {
	//可读写通道
	ch := make(chan int, 0) //创建一个int型的通道
	go TongxvnF1(ch)
	go TongxvnF2(ch)
	

}

// 接收一个只写通道
func TongxvnF1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i //写入0~99
	}
}

// 接收一个只读通道
func TongxvnF2(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}