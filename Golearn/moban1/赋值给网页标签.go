package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	// 1. 解析模板
	t, err := template.ParseFiles("./view/1.tmpl")
	if err != nil {
		fmt.Println("template parsefile failed, err:", err)
		return
	}
	// 2.渲染模板
	name := "我爱Go语言"
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", helloHandleFunc)
	http.ListenAndServe(":8086", nil)
}
