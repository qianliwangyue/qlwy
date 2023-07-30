package main

import (
	"fmt"
	"math/rand"
	"sync"
	//"sync"
)

var count int
var rw sync.RWMutex //定义一个读写互斥锁对象
func main() {
	ch := make(chan struct{}, 6)
	for i := 0; i < 3; i++ {
		go ReadCount(i, ch)
	}
	for i := 0; i < 3; i++ {
		go WriteCount(i, ch)
	}
	for i := 0; i < 6; i++ {
		<-ch //将0~5输入通道中
	}
}
func ReadCount(n int, ch chan struct{}) {
	rw.RLock() //1
	fmt.Printf("goroutine %d 进入读操作...\n", n)
	v := count
	fmt.Printf("goroutine %d 读取结束,值为：%d\n", n, v)
	rw.RUnlock() //1
	//两个1之间内容是一个整体，（完整的输入输出）
	ch <- struct{}{}
}
func WriteCount(n int, ch chan struct{}) {
	rw.Lock() //2
	fmt.Printf("goroutine %d 进入写操作...\n", n)
	v := rand.Intn(10) //获取一个0~10之间的随机数
	count = v
	fmt.Printf("goroutine %d 写入结束，新值为：%d\n", n, v)
	rw.Unlock() //2
	//两个2之间内容是一个整体，（完整的输入输出）
	ch <- struct{}{}
}
