/*类型转换用于将一个接口类型的值转换为另一个接口类型，其语法为：
T(value)
T 是目标接口类型，value 是要转换的值。

在类型转换中，我们必须保证要转换的值和目标接口类型之间是兼容的，否则编译器会报错。
*/

package main

import "fmt"

//定义一个接口
type Writer interface {
	Write([]byte) (int, error) //方法
}

//定义结构体
type StringWriter struct {
	str string
}

//接口实现方法
func (sw *StringWriter) Write(data []byte) (int, error) {
	sw.str += string(data)
	return len(data), nil
}

func main() {
	var w Writer = &StringWriter{}
	sw := w.(*StringWriter)
	sw.str = "Hello, World"
	fmt.Println(sw.str)
}

/*
以上实例中，我们定义了一个 Writer 接口和一个实现了该接口的结构体 StringWriter。
然后，我们将 StringWriter 类型的指针赋值给 Writer 接口类型的变量 w。
接着，我们使用类型转换将 w 转换为 StringWriter 类型，并将转换后的值赋值给变量 sw。
最后，我们使用 sw 访问 StringWriter 结构体中的字段 str，并打印出它的值。
*/
