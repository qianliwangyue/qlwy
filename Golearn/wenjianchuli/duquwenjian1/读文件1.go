package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./wenjianchuli/duquwenjian1/憨憨.txt")
	if err != nil {
		fmt.Printf("打开文件出错！%v\n", err)
	}
	defer file.Close()
	reader := bufio.NewReader(file) //创建reader对象来读取文件内容
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF { //end of file
			break
		}
		fmt.Print(line)
	}
}
