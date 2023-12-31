package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("form.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm() //语法分析
		fmt.Fprintln(w, "表单键值对和URL键值对：", r.Form)
		fmt.Fprintln(w, "表单键值对：", r.PostForm)
	}
}

func multiProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./view/form.html")
		t.Execute(w, nil)
	} else {
		r.ParseMultipartForm(1024)                 //从表单里提取多少字节的数据
		fmt.Fprintln(w, "表单键值对:", r.MultipartForm) //multipartform是包含2个映射的结构
	}
}

// 上传
func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./view/upload.html")//传送html
		t.Execute(w, nil)
	} else {
		r.ParseMultipartForm(4096)
		fileHeader := r.MultipartForm.File["uploaded"][0] //获取名为"uploaded"的第一个文件头

		file, err := fileHeader.Open() //获取文件
		if err != nil {
			fmt.Println("error")
			return
		}
		data, err := ioutil.ReadAll(file) //读取文件
		if err != nil {
			fmt.Println("error!")
			return
		}
		fmt.Fprintln(w, string(data))
	}
}

func main() {
	http.HandleFunc("/", upload)        //设置首页的路由
	http.HandleFunc("/1", multiProcess) //设置首页的路由
	http.HandleFunc("/2", process)      //设置首页的路由
	http.ListenAndServe(":8089", nil)   //设置监听的端口
}
