package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "道阻且长，行则降至！")

}
func haha(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "天下谁人不识君！")
}
func main() {
	mux := http.NewServeMux() //定义一个多路复用器
	mux.HandleFunc("/", hi)   //给复用器添加路径和对应处理方法
	mux.HandleFunc("/66", haha)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second, //读超时5秒
		WriteTimeout: 5 * time.Second, //写超时5秒
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
