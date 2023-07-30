package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func func1(ch1 chan<- int) { //ch1作为通道接收int
	defer wg.Done() //将等待组计数器减1
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(ch1) //关闭通道
}
func func2(ch1 <-chan int, ch2 chan<- int) { //定义ch1用来接收int通道的数据，ch2作为通道接收int
	defer wg.Done() //将等待组计数器减1
	for {           //（存在数据时）读取通道数据并乘以2赋值给ch2
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- 2 * x
	}
	once.Do(func() {
		close(ch2) //确保只关闭ch2一次
	})
}

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	wg.Add(3)
	go func1(ch1)
	go func2(ch1, ch2)
	go func2(ch1, ch2)
	wg.Wait() // 等待阻塞，直到等待组计数器为零。
	for ret := range ch2 {
		fmt.Println(ret)
	}
}

//0，2，4，6....18
//写了3个goroutine，实际只执行了1次
