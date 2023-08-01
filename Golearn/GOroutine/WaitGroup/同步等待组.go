package main

import (
	"log"
	"sync"
)

// WaitGroup的作用：（即使主程序运行结束也不会终止，直到确保协程运行结束，才退出程序）
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Println(i)
		}(i)
	}
	wg.Wait()
	log.Println("done")

}
