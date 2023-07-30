package main

import "fmt"

func Sum(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum
}

func main() {
	s := []int{6, 7, 8, 9, 10, 11}
	ch := make(chan int)
	go Sum(s, ch)
	a := <-ch
	fmt.Println(a)

}
