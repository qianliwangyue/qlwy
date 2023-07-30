package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	//err:=os.Mkdir("test",0777)//在根目录下创建新目录test：即在GoWeb中创建test文件夹
	err := os.MkdirAll("test/ff/gg", 0777) //创建多级子目录	GoWeb文件夹下创建：test/ff/gg文件夹
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 5)
	/*
		er := os.Remove("test/ff/gg") //只能删除gg,多次逐级删除
		if er != nil {
			fmt.Println(er)
		}
	*/
	err1 := os.RemoveAll("test") //删除test及旗下所有目录（删除多级目录）
	if err1 != nil {
		fmt.Println(err1)
	}

}
