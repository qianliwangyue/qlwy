package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var (
	count int32
	limit = make(chan struct{}, 10) //通过管道的大小，限制数据库并发数量（结构体管道可以接受任何数据）
)

func readDB() string {
	atomic.AddInt32(&count, 1) //计数器每次增1
	fmt.Printf("readDB()调用并发度%d\n", atomic.LoadInt32(&count))
	time.Sleep(200 * time.Millisecond)
	atomic.AddInt32(&count, -1) //自减1
	return "ok"
}
func handler1() {
	limit <- struct{}{}
	readDB() //将go rountinue放到管道中，管道容量10,所以并发限制到10以内
	<-limit
}
func main() {
	for i := 0; i < 100; i++ {
		go handler1()
	}
	time.Sleep(3 * time.Second)
}
