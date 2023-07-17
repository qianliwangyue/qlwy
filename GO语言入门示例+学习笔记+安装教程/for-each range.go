package main

import "fmt"

func main() {
	strings := []string{"google", "runoob"}
	for i, s := range strings { //字符串组strings为for循环的参数
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers { //数组number为for循环的参数
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}

/*
For-each range 循环
这种格式的循环可以对字符串、数组、切片等进行迭代输出元素。
结果：
0 google
1 runoob
第 0 位 x 的值 = 1
第 1 位 x 的值 = 2
第 2 位 x 的值 = 3
第 3 位 x 的值 = 5
第 4 位 x 的值 = 0
第 5 位 x 的值 = 0
*/
