package main

import (
	"fmt"
	"net/http"
)

// 函数
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "欢迎来到我的世界！")
}

func main() {
	server := &http.Server{
		Addr: "127.0.0.1:80", //服务器端口
	}
	http.HandleFunc("/", hello) //服务器地址，对应函数
	server.ListenAndServe()     //设置监听

}
