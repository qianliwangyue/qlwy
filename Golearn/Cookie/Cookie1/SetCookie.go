package main

import (
	"fmt"
	"net/http"
)

func testHandle(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("test_cookie")
	fmt.Printf("cookie:%#v,err:%v\n", c, err)
	cookie := &http.Cookie{
		Name:   "test_cookie",
		Value:  "lkjflksfjkljoiheofnen",
		MaxAge: 3600,
		Domain: "localhost",
		Path:   "/",
	}
	http.SetCookie(w, cookie)
	//应在具体返回之前设置cookie，否则cookie设置不成功！
	w.Write([]byte("hello")) //页面返回hello
}
func main() {
	http.HandleFunc("/", testHandle)
	http.ListenAndServe(":8085", nil)
}
