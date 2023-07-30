package main

import (
	"fmt"
)

func main() {
	var a interface{} //接口可以接受任意类型数据
	name := []string{"章所", "得到", "飞机可"}
	a = name
	fmt.Println(a)
	a = "坎坷坷"
	fmt.Println(a)
	value, ok := a.(string)
	if ok {
		fmt.Println("string")
		fmt.Println(value)
	}
	//==============================
	array := make([]interface{}, 3)
	array[0] = 1
	array[1] = "hello"
	array[2] = true
	for _, value := range array {

		switch v := value.(type) {
		case int:
			fmt.Println("int:", v)
		case string:
			fmt.Println("string:", v)
		case bool:
			fmt.Println("bool:", v)
		default:
			fmt.Println("不合理的数据类型:", v)
		}
	}
}

//Go语言里面有一个语法，可以直接判断是否是该类型的变量：value，ok=element.（T），
// 这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
// 如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false。
