package main

import (
	"fmt"
	"runtime"
	"time"
)

var num int = 1

func main() {

	for i := 1; i <= 100; i++ {
		go runtimes(1)
		num++
	}
	runtime.GOMAXPROCS(16)        //设置最大核心数
	fmt.Println(runtime.NumCPU()) //输出cpu核数

}
func runtimes(times int) int {

	for i := 1; i <= times; i++ {
		fmt.Println("runtimes", i, "次:", "你好，小屁孩~", times-i)
		time.Sleep(time.Second) //休眠1秒
		fmt.Println(num)
	}

	return times
}
