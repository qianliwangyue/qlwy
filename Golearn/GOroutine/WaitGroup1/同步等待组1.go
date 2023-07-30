package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1) //添加计数1
	go func() {
		defer wg.Done() //计数减一
		fmt.Println("1 goroutine sleep...")
		time.Sleep(2000) //休眠2秒
		fmt.Println("1 goroutine exit...")
	}()
	wg.Add(1) //添加计数1
	go func() {
		defer wg.Done() //计数减一
		fmt.Println("2 goroutine sleep...")
		time.Sleep(4000)
		fmt.Println("2 goroutine exit...")
	}()
	fmt.Println("Waiting for all goroutine")
	wg.Wait() //计数等待到0
	fmt.Println("All goroutines finished!")
}

//每个goroutine中同步进行
/*
Waiting for all goroutine
1 goroutine sleep...
2 goroutine sleep...
2 goroutine exit...
1 goroutine exit...
All goroutines finished!
*/
