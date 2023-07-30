package main

import (
	"fmt"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "欢迎来到首页，处理器：indexHandler！")
}
func hiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "欢迎来到首页，处理器：hiHandler！")
}
func webHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "欢迎来到首页，处理器：webHandler！")
}

type haihai struct { //1.结构体
	Old string
}

// ServeHTTP implements http.Handler
func (a haihai) ServeHTTP(w http.ResponseWriter, r *http.Request) { //2.结构体处理器函数
	fmt.Fprintln(w, "hi", a.Old)
}

func main() {
	mux := http.NewServeMux()         //定义一个多路复用器
	mux.HandleFunc("/", indexHandler) //给复用器添加路径和对应处理方法
	mux.HandleFunc("/hi", hiHandler)  //注册处理器函数
	mux.HandleFunc("/hi/web", webHandler)
	mux.Handle("/hi/web/hhh", haihai{Old: "Golang"}) //注册处理器（结构体对象）
	server := &http.Server{
		Addr:    ":8086",
		Handler: mux,
		//ReadTimeout:  5 * time.Second, //读超时5秒
		//WriteTimeout: 5 * time.Second, //写超时5秒
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
