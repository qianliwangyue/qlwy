package main

import (
	"fmt"
)

func foo(i int) chan int {
	ch := make(chan int)
	go func() { ch <- i }()
	return ch
}

func main() {
	ch1, ch2, ch3 := foo(3), foo(6), foo(9)

	ch := make(chan int)

	// 开一个goroutine监视各个通道数据输出并收集数据到通道ch
	go func() {
		for {
			// 监视ch1, ch2, ch3的流出，并全部流入通道ch
			select {
			case v1 := <-ch1:
				ch <- v1
			case v2 := <-ch2:
				ch <- v2
			case v3 := <-ch3:
				ch <- v3
			}
		}
	}()

	// 阻塞主线，取出通道ch的数据
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}

}
