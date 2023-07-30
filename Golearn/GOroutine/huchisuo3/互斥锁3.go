package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex //定义一个读写互斥锁对象(全局变量)
func main() {
	m = new(sync.RWMutex)
	go Reading(1)
	go Reading(2)
	time.Sleep(2 * time.Second)
}

func Reading(i int) {
	println(i, "reading strat")
	m.RLock() //1
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m.RUnlock() //1
	//两个1之间的内容，上锁后成为一个整体，即：reading start、reading一起输出，
	println(i, "reading over")
}
