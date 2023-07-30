package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 这里可以是比较耗时的事情，比如计算
func doCompute(x int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond) //模拟计算,Duration(表示持续int时间)
	return 1 + x                                                // 假如1 + x是一个很费时的计算
} //该函数即：将1+x模拟为复杂数运算（计算需要一些时间才能完成）

// 每个分支开出一个goroutine做计算，并把计算结果发送到各自通道
func branch(x int) chan int {
	ch := make(chan int)
	go func() {
		ch <- doCompute(x) //x+1
	}()
	return ch
}

/*				var chs []chan int*/
func Recombination(chs ...chan int) chan int {
	ch := make(chan int)

	for _, c := range chs {
		// 注意此处明确传值
		go func(c chan int) { ch <- <-c }(c) // 复合,即：c传给通道c，接收通道传来的c再传给ch
	}

	return ch
}

func main() {
	//返回复合结果
	result := Recombination(branch(10), branch(20), branch(30)) //即：Recombination(11,21,31)

	for i := 0; i < 3; i++ {
		fmt.Println(<-result)
	}
}
