package main

import (
	"fmt"
	"strconv"
)

func jinzita() {
	var k int
	fmt.Println("请输入要打印的金字塔层数:")
	fmt.Scanf("%d", &k)
	for i := 1; i <= k; i++ {
		for j := 1; j <= k-i; j++ {
			fmt.Print(" ")
		}

		for l := 1; l <= 2*i-1; l++ {
			if l == 1 || l == 2*i-1 || i == k {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
func main() {
	jinzita()
	str := "fduiofhoiu"
	str3 := 1515616
	str4 := strconv.Itoa(str3) //数字转换成字符串
	r := []rune(str)           //字符串转换成ascii码的数组
	fmt.Printf("%c", r[3])
	fmt.Println(str4)
	fmt.Printf("%s", []byte(str))
}
