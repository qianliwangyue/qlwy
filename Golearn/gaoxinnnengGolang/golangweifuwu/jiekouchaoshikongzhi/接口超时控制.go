package main

import (
	"net/http"
	"time"
)

func readDB() string {
	time.Sleep(400 * time.Millisecond) //假设读取数据库用时400毫秒
	return "ok"
}
func home(w http.ResponseWriter, r *http.Request) {
	var resp string
	c := make(chan struct{}, 1) //建立一个通道，容量1
	go func() {
		resp = readDB()
		c <- struct{}{} //将resp传入管道
	}()
	select {
	case <-c: //如果300ms前接收到管道数据，则resp=“ok”
	case <-time.After(300 * time.Millisecond): //如果300ms后接收到管道数据则time out
		resp = "time out!!!"
	}
	w.Write([]byte(resp)) //read数据库300ms内完成，网页显示“ok”，超出300ms，则显示"time out!!!"
}
func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8089", nil)
}
