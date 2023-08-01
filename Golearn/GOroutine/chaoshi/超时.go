package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go func() {
		time.Sleep(time.Second * 5) //执行程序所需时间
		done <- true
	}()
	select {
	case <-done: //如果3秒内接收到true，则打印done
		fmt.Println("done")
	case <-time.After(time.Second * 3):
		fmt.Println("timeout...") //3秒后，任未接收到true，则超时
	}
}
