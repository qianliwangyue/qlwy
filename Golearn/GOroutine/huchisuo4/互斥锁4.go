// 带有go关键字的函数会被Go编译器视为并发执行的函数，
// 而不带go关键字的函数则会被Go编译器视为普通的函数，会按照正常的顺序执行。
package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex //定义一个读写互斥锁对象(全局变量)
func main() {
	m = new(sync.RWMutex)
	go Read(2)
	go Write(1)
	go Write(3)
	//读写互斥，所以在写开始后，读必须等写进行完后才能继续
	time.Sleep(3 * time.Second)
}

func Read(i int) {
	println(i, "reading start")
	m.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m.RUnlock()
	println(i, "reading over")
}
func Write(i int) {
	println(i, "writing start")
	m.Lock()
	println(i, "writing")
	time.Sleep(1 * time.Second)
	m.Unlock()
	println(i, "writing over")
}
