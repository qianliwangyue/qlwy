package main

import (
	"fmt"
	"os"
)

func main() {
	//创建文件
	//Create()函数会根据传入的文件名创建文件，默认权限0666
	fp, err := os.Create("./憨憨.text") //如果文件已存在，则将文件清空######即：GoWeb/憨憨.text
	fmt.Println(fp, err)
	fmt.Printf("%T", fp) //*os.File文件的指针类型

	if err != nil {
		fmt.Println("文件创建失败！")
		//失败原因：1.文件路径不存在	2.权限不足	3.打开文件数量超过上限	4.磁盘空间不足等
		return
	}
	defer fp.Close() //关闭文件，释放资源
}
