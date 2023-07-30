package main

import (
	"fmt"
)

func main() {

	fmt.Println(getNumber())
	//两种情况：1.go func已经完成赋值，返回6
	//2.go func还未来得及赋值，返回默认值0

}
func getNumber() int {
	var i int
	go func() { //此时主程序运行getNumber()函数要得到数，而go runc()在单独给程序赋值6
		i = 6
	}()
	return i
}
