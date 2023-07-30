package main

import (
	"fmt"
	"time"
)

func Timer(Long time.Duration) chan bool {
	ch := make(chan bool)
	go func() {
		time.Sleep(Long)
		ch <- true
	}()
	return ch
}
func main() {
	timeout := Timer(5 * time.Second)

	if <-timeout {
		fmt.Println("already 5s!")
		return
	}

}
