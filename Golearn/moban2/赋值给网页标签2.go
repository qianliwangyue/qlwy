package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 定义一个UserInfo结构体
type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./view/hello.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	user := UserInfo{
		Name:   "张三",
		Gender: "男",
		Age:    28,
	}
	tmpl.Execute(w, user)
	//fmt.Println(tmpl)
}

func main() {
	http.HandleFunc("/", sayhello)
	http.ListenAndServe(":8087", nil)

}

/*
if 语句：
{{ if .condition }} {{ else }} {{ end }}
范围块：
{{ range .Items }} {{ end }}
范围块对提供的列表进行迭代。
eg:
<h1>Go templates</h1>
<p>The user is {{ .Name }}</p>
<h2>Skills:</h2>
{{ range .Skills }}
    <p>{{ . }}</p>
{{ end }}
*/
//{{.}}表示输出整个结构体内所有属性的值
//{{.Name}}表示输出该结构体内的Name的值
