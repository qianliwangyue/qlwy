package main

import (
	"fmt"
)

type Address struct {
	Name  string
	Phone int
}

func main() {
	userinfo := make(map[string]interface{})
	userinfo["username"] = "章三"
	userinfo["age"] = 20
	userinfo["hobby"] = []string{"吃", "玩"}
	fmt.Println(userinfo["username"])
	fmt.Println(userinfo["hobby"])
	//fmt.Println(userinfo["hobby"][1])错误：接口不支持索引
	addr := Address{
		Name:  "里斯",
		Phone: 19992018470,
	}
	userinfo["address"] = addr
	//fmt.Println(userinfo["address"].Name)错误：接口访问不了结构体内成员
	//================================================================================
	//下面是断言：用来判断接口内储存的类型并获取该类型的值
	hobby, ok := userinfo["hobby"].([]string) //判断接口内存放的类型是否为string切片
	//如果断言为真，则赋值给hobby；断言为假，则不赋值。
	if ok {
		fmt.Println(hobby[1]) //通过hobby就可以访问切片内容（索引）
	}
	ad, o := userinfo["address"].(Address) //判断接口内存放的类型是否为:结构体Address
	if o {
		fmt.Println(ad.Phone)
	}
}
