package main

import (
	"fmt"
	"net/http"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	// 设置一个 301 重定向,页面跳转
	w.Header().Set("Location", "https://www.shirdon.com")
	w.WriteHeader(301)
}

func main() {
	http.HandleFunc("/redirect", Redirect)
	err := http.ListenAndServe(":8086", nil)
	if err != nil {
		fmt.Println(err)
	}
}
