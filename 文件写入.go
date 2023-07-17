package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	Writewj("C:/Users/Administrator/Desktop/2.txt", "\n\n\n我喜欢徐欣蕾\n")
	Wwj("C:/Users/Administrator/Desktop/2.txt", "\n但也许再也见不到了~\n")
	//Write("C:/Users/Administrator/Desktop/2.txt", "\n\n再见容易，再见难！~\n")
}
func Writewj(a string, b string) {
	//不存在该文件时，创建它|读写权限|只写|追加到末尾
	con, err := os.OpenFile(a, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer con.Close() //defer：延迟，即稍后自动关闭文件
	/*
		0777：-rwxrwxrwx，创建了一个普通文件，
		所有人拥有所有的读、写、执行权限
		0666：-rw-rw-rw-，创建了一个普通文件，
		所有人拥有对该文件的读、写权限，但是都不可执行
		0644：-rw-r--r--，创建了一个普通文件，
		文件所有者对该文件有读写权限，用户组和其他人只有读权限，没有执行权限
	*/
	if err != nil {
		fmt.Printf("写入失败%v", err)
		return
	}

	//con.Write([]byte(b))//写入字符组b
	con.WriteString(b) //写入字符串b

}
func Wwj(a string, b string) {
	con, err := os.OpenFile(a, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	defer con.Close() //defer：延迟，即稍后自动关闭文件
	if err != nil {
		fmt.Printf("写入失败%v", err)
		return
	}
	con2 := bufio.NewWriter(con)
	con2.WriteString(b) //将数据先写入缓存区
	con2.Flush()        //从缓存区写入文件

}

// 清空原文件，写入字符串b
func Write(a string, b string) {
	err := ioutil.WriteFile(a, []byte(b), 0644)
	if err != nil {
		fmt.Printf("写入失败%v", err)
		return
	}
}
