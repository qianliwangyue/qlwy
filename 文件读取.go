package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	Readline("C:/Users/Administrator/Desktop/2.txt")
	fmt.Println("==============================================")
	Read_all("C:/Users/Administrator/Desktop/2.txt")
}

func Readline(a string) {
	read, err := os.Open(a)
	if err != nil && err != io.EOF {
		fmt.Printf("读取失败：%v\n", err)
	}
	read1 := bufio.NewReader(read)
	for {
		line, err := read1.ReadString('\n') //循环每次读取一行
		if err != nil && err != io.EOF {
			fmt.Printf("读取失败：%v\n", err) //每次读取一行后，输出一行
			return
		}
		fmt.Print(line)
		if err == io.EOF { //读完结束循环
			fmt.Printf("\n\n读取完毕---------\n\n")
			break
		}
	}
	read.Close() //关闭文件
}

func Read_all(a string) {

	wenjian, err := ioutil.ReadFile(a) //打开文件

	if err != nil {
		fmt.Printf("读取失败：%v\n", err)
	}
	fmt.Println(string(wenjian))
}
