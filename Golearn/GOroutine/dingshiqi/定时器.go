package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Tick(time.Second)
	t2 := time.Tick(time.Second * 3)
	t3 := time.Tick(time.Second * 5)
	for {
		select {
		case <-t1:
			fmt.Println("每1秒执行一次")
		case <-t2:
			fmt.Println("每3秒执行一次")
		case <-t3:
			fmt.Println("每5秒执行一次")
		}

	}
}
