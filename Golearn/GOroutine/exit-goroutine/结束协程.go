package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		go func() {
			fmt.Println("go协程内部函数")
			//runtime.Goexit() //退出子，当前协程
			//os.Exit(-1)//崩溃，如果协程比主函数运行快，那么部分来不及运行的主函数内容将不运行。

		}()
		fmt.Println("子协程结束")
	}()

	//======================================================
	fmt.Println("go主进程")
	time.Sleep(3 * time.Second)
	fmt.Println("over")
}
