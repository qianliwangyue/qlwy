package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Print("err", err)
	}
	closer := res.Body                 //url的body赋值给closer
	bytes, _ := ioutil.ReadAll(closer) //err舍弃，读取url的body
	fmt.Println(string(bytes))
	//获取百度首页的HTML文档

}
