package main

import (
	"fmt"
	"time"
)

func jishi() {
	start := time.Now().Unix()
	var sum = 0
	for i := 0; i < 100; i++ {
		sum += 1
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println(sum)
	end := time.Now().Unix()
	fmt.Printf("程序执行用时%v秒\n", end-start)
}
func main() {
	start1 := time.Now()
	jishi()
	r := time.Since(start1)
	fmt.Println("程序用时：", r)
}
