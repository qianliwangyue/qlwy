package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("憨憨.text")
	//打开根目录的文件|可以加上文件的路径eg:"./wenjianchuli/duxiewenjian/憨憨.txt"
	if err != nil {
		fmt.Printf("打开文件出错：%v\n", err)
	}
	fmt.Println(file)

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Printf("关闭文件出错：%v\n", err)
	}

}
