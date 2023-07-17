package main

import "fmt"

var hhu int = 6 //全局变量

var (
	j1 = 65
	j2 = "dd"
)

func main() {
	var age uint //定义变量age：无符号整型
	age = 18     //赋值
	age = 16
	var name string = "fewf" //定义变量同时赋值
	var num float32 = 3.015
	var num1 = int(num) //强制类型转换
	var s = 5156        //未定义类型，则系统自动推理
	sex := "男"          //省略var

	fmt.Println(sex)
	fmt.Println(s)
	fmt.Println("age=", age)
	fmt.Println(name)
	fmt.Println(num1)
	fmt.Println("===============================")
	//声明多个变量
	var n1, n2, n3 int             //一次声明多个变量未赋值
	var s1, s2, s3 = 1, "dd", "65" //一次声明多个变量并赋值
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println("全局变量1：", j1)
	fmt.Println("全局变量2：", j2)

	fmt.Println("===============================")
	var f float64
	var b bool
	var k string
	fmt.Printf(" %v %v %q\n", f, b, k) //  0 false ""
	//%v :值  %q :输出双引号内放字符串值

	fmt.Println("===============================")

	//va := 1

	var v1 uint8 = 1
	//va = 1

	g := "Runoob" //定义时未声明变量类型，则输出也省略%输出格式

	fmt.Println(g)
	fmt.Printf("%d", v1)
}
