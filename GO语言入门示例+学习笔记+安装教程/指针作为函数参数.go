/*
Go 语言允许向函数传递指针，只需要在函数定义的参数上设置为指针类型即可。
以下实例演示了如何向函数传递指针，并在函数调用后修改函数内的值，：
*/
package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值 : %d\n", a)
	fmt.Printf("交换前 b 的值 : %d\n", b)
	fmt.Println("==================================================")
	/* 调用函数用于交换值
	 * &a 指向 a 变量的地址
	 * &b 指向 b 变量的地址
	 */
	swap(&a, &b) //调用函数
	fmt.Println("==================================================")
	swap2(&a, &b) //调用函数
}

//定义函数
func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
	fmt.Printf("交换后的值 : %d\n", *x)
	fmt.Printf("交换后的值 : %d\n", *y)
}

//定义函数
func swap2(x *int, y *int) {
	*x, *y = *y, *x
	fmt.Printf("交换后的值 : %d\n", *x)
	fmt.Printf("交换后的值 : %d\n", *y)
}
