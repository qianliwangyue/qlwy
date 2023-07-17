package main

import "fmt"

// 猴子每天吃一半的桃子，再多吃一个，吃到第10天（还没吃）剩余1个桃子，求第1天有几个桃子？
// 提示：第9天的桃子数=（第10天+1）*2
func A(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("输入天数不对！")
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (A(n+1) + 1) * 2
	}

}
func main() {
	fmt.Println(A(1))
}
