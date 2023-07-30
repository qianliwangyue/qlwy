package main

import (
	"errors"
	"fmt"
)

// 捕获异常: defer 、panic 、recover
// panic打印错误并立即结束程序
func test() {
	defer func() {

		if err := recover(); err != nil {
			fmt.Println("err:", err)
			//panic(err)
		}
		//等价于:
		//err:=recover()
		//if err!=nil{fmt.Println("err:", err)}
	}()
	n1 := 10
	n2 := 0
	r := n1 / n2 //==>10/0出错
	fmt.Println(r)
}
func readconf(name string) (err error) {
	if name == "config.txt" {
		//读取
		return nil
	} else {
		return errors.New("读取文件错误！")
	}
}
func main() {
	test()
	fmt.Println("下面的程序")
	r := readconf("222ddwf")
	fmt.Println("error:", r)
}
