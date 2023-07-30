package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./wenjianchuli/duquwenjian1/憨憨.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("文件打开失败：%v\n", err)
	}
	defer file.Close()

	context := []byte("你好，写入成功666！")
	if _, err = file.Write(context); err != nil {
		fmt.Printf("写入失败：%v\n", err)
	}
	fmt.Println("写入成功！")
}
