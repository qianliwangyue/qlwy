package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	conn, err := rpc.Dial("tcp", "127.0.0.1:6000") //连接服务器端
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	defer conn.Close()
	var reply string
	err = conn.Call("hello.Helloworld", "李白", &reply)
	if err != nil {
		fmt.Println("call err:", err)
	}
	fmt.Println(reply)
}
