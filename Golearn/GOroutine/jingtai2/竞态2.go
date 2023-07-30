package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["a"] = "one" //1
		c <- true
	}()
	m["b"] = "two" //2
	<-c
	for k, v := range m {
		fmt.Println(k, v) //存在竞态，可能先输出1，有可能先输出2
	}
}

/*
cd GOroutine
cd jingtai2
go run -race 竞态2.go

结果：Found 1 data race(s)存在竞态
*/
