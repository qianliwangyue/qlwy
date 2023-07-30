package main

import (
	"fmt"
)

func main() {
	ch, quit := make(chan int), make(chan int)
	go func() {
		ch <- 8
		quit <- 1
	}()
	for isQuit := false; !isQuit; { //for循环省略第三个参数
		select {
		case v := <-ch:
			fmt.Printf("receive %d from ch", v) //8
		case <-quit:
			isQuit = true
		}
	}
}
