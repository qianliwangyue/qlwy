package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var i int64 = 0
var wg sync.WaitGroup
var lock sync.Mutex

func incre() { //普通自增
	i += 1
	wg.Done()
}
func muteincre() { //加锁自增
	lock.Lock()
	defer lock.Unlock()
	i += 1
	wg.Done()

}
func atoincre() {
	atomic.AddInt64(&i, 1) //原子操作：作为一个整体执行，只有2种状态（全部未执行，全部已执行，没有执行中的状态），比加锁消耗资源小
	wg.Done()
}
func main() {
	for j := 0; j < 100000; j++ {
		wg.Add(1)
		//go incre()     //不到100000
		//go muteincre() //100000
		go atoincre() //原子操作//100000
	}
	wg.Wait()
	fmt.Println(i)
}
