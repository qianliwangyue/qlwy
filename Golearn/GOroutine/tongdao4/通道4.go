package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 5; i < 9; i++ {
			ch <- i
		}
	}()
	for da := range ch {
		fmt.Println(da)
		if da == 8 {
			break
		}

	}

}
