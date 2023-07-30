package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	Func := func(wg *sync.WaitGroup, id int) {
		defer wg.Done() //计数减一
		fmt.Printf("%v goroutine start...\n", id)
		time.Sleep(2 * time.Second)
		fmt.Printf("%v goroutine exit...\n", id)
	}
	var wg sync.WaitGroup
	const N = 3
	wg.Add(N) //计数加3
	for i := 0; i < N; i++ {
		go Func(&wg, i)
	}
	fmt.Println("Waiting for all goroutine")
	wg.Wait() //计数减到0时，再运行下一步
	fmt.Println("All goroutines finished!")
}

/*
Waiting for all goroutine
2 goroutine start...
0 goroutine start...
1 goroutine start...
0 goroutine exit...
2 goroutine exit...
1 goroutine exit...
All goroutines finished!
*/
