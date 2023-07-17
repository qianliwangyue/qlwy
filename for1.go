package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i <= 10; i++ { //别加括号容易报错go语法不严谨
		sum += i
	}
	fmt.Println(sum)
}
