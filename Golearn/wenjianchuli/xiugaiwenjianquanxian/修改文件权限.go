package main

import (
	"fmt"
	"os"
)

func main() {
	fp, err := os.Create("目录/555.txt") //默认权限0666
	if err != nil {
		fmt.Println("文件创建失败！")
	}
	defer fp.Close()

	fileInfo, err := os.Stat("目录/555.txt") //Stat返回文件信息和err
	if err != nil {
		fmt.Println("fileInfo返回失败,err:", err)
	}
	fmt.Printf("%v\t %v\t%v\t\n", fileInfo.Name(), fileInfo.Size(), fileInfo.Mode())

	fmt.Println("=============================================================")

	os.Chmod("目录/555.txt", 0777) //重新赋于权限0777

	fileInfo, err = os.Stat("目录/555.txt")
	if err != nil {
		fmt.Println("fileInfo返回失败,err:", err)
	}

	fmt.Printf("%v\t %v\t%v\t", fileInfo.Name(), fileInfo.Size(), fileInfo.Mode())

}
