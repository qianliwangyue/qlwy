package main

import (
	"fmt"
)
func main(){


ch1:= make (chan int, 2)
ch2 := make (chan int, 1)
//ch1<-6
ch1<-4
//ch2<-5

//ch3<-"哈哈"
select {
    
case <-ch1:
    fmt.Println("ch1 收到数据")
case <-ch2:
    fmt.Println("ch2 收到数据")
default:
    fmt.Println("接收失败")
}
}

//如果ch1,ch2都收到数据，则随机返回fmt.println中内容