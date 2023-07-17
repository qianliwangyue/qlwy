package main

import "fmt"

func main() {
	sum := 1
	//init 和 post 参数是可选的，我们可以直接省略它，类似 While 语句。
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum) //自身*2直到>10  即：16

}
