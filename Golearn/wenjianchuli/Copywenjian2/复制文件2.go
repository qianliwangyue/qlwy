package main

import (
	"io"
	"log"
	"os"
)

func DoCopy(srcfileName string, newfileName string) {
	srcfile, err := os.Open(srcfileName)
	if err != nil {
		log.Fatalln("原文件打开失败,err:", err)
	}
	defer func() { //闭包函数
		err = srcfile.Close()
		if err != nil {
			log.Fatalln("原文件关闭失败,err:", err)
		}
	}()

	newfile, err := os.Create(newfileName)
	if err != nil {
		log.Fatalln("newfile 创建失败,err:", err)
	}
	defer func() { //闭包函数
		err = newfile.Close()
		if err != nil {
			log.Fatalln("newfile关闭失败,err:", err)
		}
	}()

	var tmp = make([]byte, 1024*4) //定义一个临时的内存tmp（可容纳1024*4个字节数）
	for {
		n, err := srcfile.Read(tmp) //读取源文件，并储存到tmp中
		if err != nil {
			if err == io.EOF {
				log.Fatalf("读取完毕！读取了%d个字节", n)
				return
			} else {
				log.Fatalln("读取过程出错,err:", err)
			}
		}
		n, err = newfile.Write(tmp[:n]) //读取的内容写入新文件
		if err != nil {
			log.Fatalln("复制过程出错,err:", err)
		} else {
			log.Fatalf("复制成功，复制了%d个字节", n)
			break
		}

	}

}

func main() {
	err := os.MkdirAll("目录/目录2", 0777)
	if err != nil {
		log.Fatalln("目录2创建出错:", err)
	}

	fp, err := os.Create("目录/目录2/66.txt")
	if err != nil {
		log.Fatalln("src文件创建失败,err:", err)
	}

	fp.WriteString("加油！年轻人！")

	DoCopy("目录/目录2/66.txt", "目录/目录2/哈哈哈.txt")
}
