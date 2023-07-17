package main

import (
	"fmt"
)

func main() {
	var a uint8
	fmt.Println("ce" + "vv") //字符串连接，结果：cevv
	fmt.Println("请输入1个整数：")
	fmt.Scanf("%d", &a)
	if a > 0 {
		fmt.Println(a)
	}

	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d AND Date=%s"
	//Sprintf 根据格式化参数生成格式化的字符串并返回该字符串。注意：所以他的类型是字符串，而不是动词输出。
	var target_url = fmt.Sprintf(url, stockcode, enddate) //结果：Code=123 AND Date=2020-12-31
	//等价于：fmt.Sprintf("Code=%d AND Date=%s", stockcode, enddate)
	fmt.Println(target_url)

	fmt.Println("-----------------------------------------------")

	var stockcode1 = 123
	var enddate1 = "2020-12-31"
	var url1 = "Code1=%d and Date=%s"
	//Printf 根据格式化参数生成格式化的字符串并写入标准输出。
	fmt.Printf(url1, stockcode1, enddate1) //结果：同上，但没有换行

	fmt.Println("\n-----------------------------------------------")

}
