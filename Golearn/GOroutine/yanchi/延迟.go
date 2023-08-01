package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 3)
	fmt.Println(time.Now())
	<-timer.C //延迟3秒
	fmt.Println(time.Now())
}
