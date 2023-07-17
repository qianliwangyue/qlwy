package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("---------------------------------------------------------")
	fmt.Println(time.Now().Year(), time.Now().Month(), time.Now().Day())
	fmt.Println(time.Now().Hour(), ":", time.Now().Minute(), ":", time.Now().Second())
	fmt.Println(time.Now().Local().Weekday())
	//时间戳：
	fmt.Println(time.Now().Unix())     //1970.1.1到现在的秒数
	fmt.Println(time.Now().UnixNano()) //1970.1.1到现在的毫秒数
	t := time.Unix(1682874381, 0)      //时间戳转换为当前时区时间
	fmt.Println(t)
	s := time.Duration.Hours(time.Second)
	fmt.Println(s)
	a := "1561f ff f3fg gh 的"
	fmt.Sscan(a)
	fmt.Println(a)
}
