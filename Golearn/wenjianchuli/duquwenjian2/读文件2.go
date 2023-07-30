package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//用io/ioutil.ReadFile()函数一次性将文件读取到内存中
	filepath := "./wenjianchuli/duquwenjian1/憨憨.txt"
	content, err := ioutil.ReadFile(filepath) //直接读取到内存中
	if err != nil {
		fmt.Printf("读取文件出错：%v\n", err)
	}
	fmt.Printf("%v\n", content)         //ASCII码数组
	fmt.Printf("%v\n", string(content)) //字符串
}
