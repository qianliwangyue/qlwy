package main

import "fmt"

type human struct {
	name string
	age  int
}

func (h *human) eat(thing string) {
	fmt.Printf("%s吃%s\n", h.name, thing)
}

type student struct {
	hum human //student嵌套了human
	sid string
}
type stu struct {
	human //stu继承了human
	class string
}

func main() {
	s1 := student{
		hum: human{
			name: "章三",
			age:  16,
		},
		sid: "2006150405",
	}
	fmt.Printf("%+v\n", s1)
	//=============================================
	s1.sid = "15161"
	s1.hum.age = 15
	s1.hum.name = "章四"
	fmt.Printf("%+v\n", s1)
	fmt.Println("==============================================")
	s2 := stu{
		human: human{
			name: "里二",
			age:  15,
		},
		class: "信息2006",
	}
	fmt.Printf("%+v\n", s2)
	//=============================================
	s2.age = 17
	s2.class = "信息2007"
	s2.name = "里三"
	fmt.Printf("%+v\n", s2)
	//=============================================
	s1.hum.eat("饼干") //嵌套间接调用"方法"
	s2.eat("肉")      //继承可以直接调用"父类方法"

}

// %v 按默认格式输出
// %+v 在%v的基础上额外输出字段名
// %#v 在%+v的基础上额外输出类型名
