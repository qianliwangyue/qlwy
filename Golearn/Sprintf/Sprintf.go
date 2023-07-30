package main

import (
	"fmt"
	"os"
)

func main() {
	var progress = 2
	var target = 8 // 两参数格式化
	title := fmt.Sprintf("已采集%d个药草, 还需要%d个完成任务", progress, target)
	//Sprintf（）返回一个字符串
	fmt.Println(title)
	pi := 3.14159 // 按数值本身的格式输出
	variant := fmt.Sprintf("%v %v %v", "月球基地", pi, true)
	fmt.Println(variant) // 匿名结构体声明, 并赋予初值
	profile := &struct {
		Name string
		HP   int
	}{Name: "rat", HP: 150}
	fmt.Printf("使用'%%+v' %+v\n", profile)
	fmt.Printf("使用'%%#v' %#v\n", profile)
	fmt.Printf("使用'%%T' %T\n", profile)

	fmt.Println("============================================")
	c, er := os.Create("1.txt")
	if er != nil {
		fmt.Println("create file err:", er)
	}
	n, err := fmt.Fprintf(c, "%s", "sssss") //格式化写入到文件/stdout（控制台）...
	if err != nil {
		fmt.Println("写入失败:", err)
	}
	fmt.Println("写入", n, "个字符")
	//它将格式化后的字符串输出到指定的文件中
}
