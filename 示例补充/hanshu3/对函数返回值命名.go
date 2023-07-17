package main

import "fmt"

// 返回值命名决定了返回的值的顺序sum、sub
// -------参数列表----返回值1,返回值2...
func ss(a, b int) (sum int, sub int) {
	sub = a - b
	sum = a + b
	//return sum, sub
	return
}
func main() {
	sum, sub := ss(5, 2)
	fmt.Println(sum, sub) //7 3
	fmt.Println(ss(5, 2)) //7 3
	a := 10
	b := 2
	dd(&a, &b)
	fmt.Println(a, b)

}

// 两数交换（用指针函数实现）
func dd(a, b *int) {
	t := *a
	*a = *b
	*b = t
}
