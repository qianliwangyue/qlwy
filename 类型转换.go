/*
Go 语言类型转换基本格式如下：
type_name(expression)
type_name 为类型，expression 为表达式。
*/
/*
将整型转换为浮点型：
var a int = 10
var b float64 = float64(a)
*/
/*
var str string = "10"
var num int
num, _ = strconv.Atoi(str)
以上代码将字符串变量 str 转换为整型变量 num。

注意，strconv.Atoi 函数返回两个值，第一个是转换后的整型值，
第二个是可能发生的错误，我们可以使用空白标识符 _ 来忽略这个错误
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sum int = 17
	var count int = 5
	var mean float32
	//将整型转化为浮点型
	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)

	fmt.Println("=============================================")
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		//将字符串转换为整数
		fmt.Printf("字符串 '%s' 转换为整数为：%d\n", str, num)
	}

	fmt.Println("=============================================")
	//将整数转换为字符串：
	num2 := 123
	str2 := strconv.Itoa(num2)
	fmt.Printf("整数 %d  转换为字符串为：'%s'\n", num2, str2)

	fmt.Println("=============================================")
	//将字符串转换为浮点数
	str3 := "3.14"
	num3, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Printf("字符串 '%s' 转为浮点型为：%f\n", str3, num3)
	}

	fmt.Println("=============================================")
	//将浮点数转换为字符串：
	num4 := 3.14
	str4 := strconv.FormatFloat(num4, 'f', 2, 64)
	fmt.Printf("浮点数 %f 转为字符串为：'%s'\n", num4, str4)

}
