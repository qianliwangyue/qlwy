package main

import (
	"fmt"
	"time"
)

func handtime(a, b int) string {
	t0 := time.Now()
	defer func() {
		fmt.Printf("use time :%d ms\n", time.Since(t0).Milliseconds()) //返回handtime函数运行用时
	}()
	//不能直接defer Print，不然返回用时0ms
	if a > b {
		time.Sleep(100 * time.Millisecond) //0.1秒
		return "ok"
	} else {
		time.Sleep(200 * time.Millisecond)
		return "ok"
	}
}
func main() {
	handtime(2, 3)
}
