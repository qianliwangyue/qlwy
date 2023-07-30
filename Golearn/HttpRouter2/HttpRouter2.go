package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New() //创建路由对象								注意"_"前后各一个空格
	router.GET("/default", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("default get"))
	})
	router.POST("/default", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("default post"))
	})

	//路由精确匹配														//注意p
	/*	router.GET("/user/:name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			w.Write([]byte("user name:" + p.ByName("name")))
		}) //输入url："localhost:8088/user/张三"，结果：user name:张三

	*/
	//路由匹配所有
	router.GET("/user/*name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("user name:" + p.ByName("name")))
	}) // 输入url："localhost:8088/user/张三/都无法..." ，结果：user name:/张三/都无法...

	http.ListenAndServe(":8088", router)
}
