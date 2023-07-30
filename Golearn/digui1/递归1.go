package main

import "fmt"

// 斐波那契数：1,1,2,3,5,8,13,21.....
// 题1：求第n位斐波那契数是？
func ti1(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return ti1(n-1) + ti1(n-2)
	}

}

// 题2：f(1)=3;f(n)=2*f(n-1)+1;求f(n)的值？
func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
	}

}
func main() {
	var n int
	fmt.Println("请输入斐波那契数的位数：")
	fmt.Scanf("%d", &n)
	fmt.Println("\n第", n, "位斐波那契数是:", ti1(n))

	var n2 int
	fmt.Println("请输入一个整数：")
	fmt.Scanf("%d", &n2)
	fmt.Printf("f(%d)=%d\n\n", n2, f(n2))
}
