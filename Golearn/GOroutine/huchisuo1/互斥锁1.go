package main

import (
	"log"
	"sync"
	"time"
)

//10个窗口卖票，当剩余0张票时结束程序。

var ticket = 5
var wg sync.WaitGroup
var mu sync.Mutex

func sellticket() {
	defer wg.Done()
	mu.Lock()         //加互斥锁
	defer mu.Unlock() //解互斥锁（这里是整个函数加锁，defer函数执行完成时解锁。如果需要具体某些步骤加锁去掉defer，到对应位置加锁即可）
	if ticket <= 0 {
		return
	}
	time.Sleep(100 * time.Millisecond)
	ticket -= 1
	log.Println("sell ticket -1")
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go sellticket()
	}
	wg.Wait()
	log.Println(ticket) //-5
}

//未加锁时，最后剩余票-5，因为多个协程出现了脏读、脏写（即：多个协程同时操控了该数据【而不是轮流操控】）
//加锁后，最后剩余票0
