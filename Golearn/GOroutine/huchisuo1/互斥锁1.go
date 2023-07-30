package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	hh := sync.WaitGroup{}
	mutex.Lock() //加锁1
	fmt.Println("Locked")
	for i := 0; i <= 5; i++ {
		hh.Add(1) //计数器添加1个goroutine
		go func(i int) {
			fmt.Println("Not lock:", i)
			mutex.Lock() //加锁2
			fmt.Println("Lock:", i)
			time.Sleep(time.Second)
			fmt.Println("UnLock:", i)
			mutex.Unlock()  //解锁2
			defer hh.Done() //计数器减一，即hh.add()结束
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Unlocked")
	mutex.Unlock() //解锁1
	hh.Wait()
}

//加锁与解锁之间内容是一起运行的
//即：Look：和Unlock：一起输出、剩下的Locked、Not lock:（循环5次） 、Unlocked 一起运行
