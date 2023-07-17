/*
接口类型转换有两种情况：类型断言和类型转换。
类型断言用于将接口类型转换为指定类型，其语法为
value.(type)
或者
value.(T)
其中 value 是接口类型的变量，type 或 T 是要转换成的类型。
如果类型断言成功，它将返回转换后的值和一个布尔值，表示转换是否成功。
*/
package main

import "fmt"

func main() {
	var i interface{} = "Hello, World"
	str, a := i.(string) //返回转换后的值和一个布尔值，str转换后的值，a布尔值
	if a {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("conversion failed")
	}
}
