package main

import (
	"fmt"
	"os"
)

func main() {
	cmds := os.Args //读取终端输入的内容（字符串）
	for key, cmd := range cmds {
		fmt.Println("key:", key, ",cmd:", cmd, ",cmds len:", len(cmds))
		//os.Args[0]==>程序名
		//os.Args[1]==>第一个参数
		//os.Args[2]==>第二个参数
		//.......
	}
	if len(cmds) < 2 {
		fmt.Println("请输入正确的参数")
		return
	}
	switch cmds[1] {
	case "hello":
		fmt.Println("hello")
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("default")
	}

}

//先go build switch.go
//终端输入： ./switch hello world
//运行结果：读取三个字符串
// key: 0 ,cmd: ./switch ,cmds len: 3
// key: 1 ,cmd: hello ,cmds len: 3
// key: 2 ,cmd: world ,cmds len: 3
