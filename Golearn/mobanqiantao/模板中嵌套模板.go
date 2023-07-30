package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func tmpSample(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./view/t.html", "./view/ul.html")
	//将ul.html中的内容插入到t.html的{{template "ul.html"}}中，即：两个HTML网页拼成一个HTML
	if err != nil {
		fmt.Println("create template failed,err:", err)
		return
	}

	user := UserInfo{
		Name:   "张三",
		Gender: "男",
		Age:    25,
	}
	tmpl.Execute(w, user)
}

func main() {
	http.HandleFunc("/", tmpSample)
	http.ListenAndServe(":8080", nil)

}
