package main

import (
	"fmt"
	"net"
)

func main() {
	//创建服务器端UDP地址结构：指定地址和端口号
	laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8023")
	if err != nil {
		fmt.Println("net.ResolveUDPAddr err:", err)
		return
	}
	//监听客户端连接
	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		fmt.Println("net.ListenUDP err:", err)
		return
	}
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, raddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("conn.ReadFromUDP err:", err)
			return
		}
		fmt.Printf("接收客户端[%s]:%s", raddr, string(buf[:n]))
		conn.WriteToUDP([]byte("你好，我是服务器端-"), raddr) //简单回复数据给客户端
	}

}
