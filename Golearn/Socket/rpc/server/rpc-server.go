package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type humen struct {
}

func (h *humen) Helloworld(name string, resp *string) error {
	*resp = name + "你好"
	return nil
}
func main() {
	err := rpc.RegisterName("hello", new(humen))
	if err != nil {
		fmt.Println("rpc注册失败:", err)
		return
	}
	listener, err1 := net.Listen("tcp", "127.0.0.1:6000")
	if err1 != nil {
		fmt.Println("net.Lense err:", err1)
		return
	}
	defer listener.Close()
	fmt.Println("开始监听...")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Accept err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("连接成功")
	rpc.ServeConn(conn)
}
