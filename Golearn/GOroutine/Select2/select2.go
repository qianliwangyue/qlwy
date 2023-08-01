package main

import (
	"fmt"
	"time"
)

func main() {
	ser := make(chan string, 10)
	cli := make(chan string, 10)

	//服务器端
	go func() {
		for {
			select {
			case msg, ok := <-ser:
				if !ok {
					fmt.Println("ser closed!")
					break
				}
				fmt.Printf("ser read:%s\n", msg)
				cli <- fmt.Sprintf("reply to :%s\n", msg)
			}
		}
	}()
	//客户端
	go func() {
		for {
			select {
			case msg, ok := <-cli:
				if !ok {
					fmt.Println("cli closed!")
					break
				}
				fmt.Printf("ser read:%s\n", msg)

			}
		}
	}()
	ser <- "hello ,I am  Petter."
	time.Sleep(2 * time.Second)
}
