package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex //定义互斥锁

func main() {

	runtimes(10) //go协程-关键字，不加则按顺序运行runtimes而后main
	//加go后：同时运行抢占cpu轮流输出，主线程main结束，runtimes即使没运行完也会终结
	for i := 1; i <= 10; i++ {
		fmt.Println("main", i, "次:", "你好，小屁孩~", 10-i)
		time.Sleep(time.Second * 2) //休眠2秒
	}
}
func runtimes(num int8) int8 {
	lock.Lock() //开始互斥锁
	var i int8
	for i = 1; i <= num; i++ {
		fmt.Println("runtimes", i, "次:", "你好，小屁孩~", num-i)
		time.Sleep(time.Second) //休眠1秒
	}
	lock.Unlock() //解除锁
	return num

}
