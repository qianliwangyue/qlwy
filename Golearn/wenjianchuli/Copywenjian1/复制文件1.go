package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("目录/哈哈.zip")
	if err != nil {
		fmt.Println("文件创建失败：", err)
	}
	srcfile, err := os.Open("目录/哈哈.zip")
	if err != nil {
		fmt.Println("文件打开失败：", err)
		return
	}
	file.WriteString("不要辜负最好的自己！") //30个btye，一个中文字符占个字符
	defer srcfile.Close()

	newfile, err := os.OpenFile("目录/new哈哈.zip", os.O_WRONLY|os.O_CREATE, 0775)
	if err != nil {
		fmt.Println("newfile create failed ,err:", err)
		return
	}
	defer newfile.Close()

	result, err := io.Copy(newfile, srcfile)
	if err != nil {
		fmt.Println("Copy failed,err:", err)
	} else {
		fmt.Println("Copy file succeed ,types:", result)
	}

}
