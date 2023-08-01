package main

import (
	"fmt"
	"sync"
)

// 直接map并发会脏读或脏写，需要用map锁
// ===============================================================================
// var m = make(map[int]int)
var wg sync.WaitGroup
var sm sync.Map

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			//m[i] = i
			//fmt.Println(m[i])
			sm.Store(i, i) //给键赋值
			wg.Done()
		}()
		wg.Wait()
	}
	sm.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
