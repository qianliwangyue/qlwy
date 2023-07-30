package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./open.text", os.O_CREATE|os.O_APPEND, 0666) //只读方式打开文件
	//不存在该文件则自动创建
	if err != nil {
		fmt.Printf("打开文件失败！")
		return
	}
	file.Close()
	 //前面要是+defer 则后面文件删除会失败。因为defer表示程序运行结束后才关闭文件，所以文件打开状态下删除不了！！！

	err1 := os.Remove("./open.text")
	if err1 != nil {
		fmt.Println("删除文件失败！")
		return
	}

}
