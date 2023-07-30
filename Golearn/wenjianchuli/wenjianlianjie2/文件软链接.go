package main

import (
	"fmt"
	"os"
)

func main() {
	fp, err := os.Create("目录/3333.txt")
	if err != nil {
		fmt.Println("文件创建失败!,err:", err)
	}
	defer fp.Close()
	err = os.Symlink("目录/3333.txt", "目录/4444.txt")
	//软链接相当于创建快捷方式，而硬链接相当于建立文件副本备份原文件（两个相同的文件）
	if err != nil {
		fmt.Println("Link失败,err:", err)
	}
}
