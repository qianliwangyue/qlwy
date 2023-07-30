package main

import (
	"fmt"
)

// 闭包即匿名函数的特点：可以在主函数中定义,直接赋值运行
func main() {
	ch := make(chan string)
	go func(a int) {
		fmt.Println("开始goroutine")
		ch <- "好难啊"
		fmt.Println("退出goroutine")
		fmt.Println(a)//6
	}(6)
/*
	go func() {
		fmt.Println("开始goroutine")
		ch <- "好难啊"
		fmt.Println("退出goroutine")
	}()
*/
	fmt.Println("等待goroutine")
	d := <-ch
	fmt.Println("完成")
	fmt.Println("goroutine:", d)

}

/*
func hh(ch chan string){
	fmt.Println("开始goroutine")
	ch<-"好难啊"
	fmt.Println("退出goroutine")
}
func main(){
	ch:=make(chan string)
	go hh(ch)
	fmt.Println("等待goroutine")
	d:=<-ch
	fmt.Println("完成")
	fmt.Println("goroutine:",d)
}
*/
