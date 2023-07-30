package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./wenjianchuli/duquwenjian1/writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("我真是太难了！")                     //第一种写入方法
	n, err := file.WriteAt([]byte("\n嗨，不容易啊！"), 48) //第二种写入方法
	if err != nil {
		panic(err)
	}
	fmt.Println(n) //第二种写入方法,返回写入的字节数
}
