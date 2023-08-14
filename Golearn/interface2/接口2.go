package main

import "fmt"

// 接口中函数可以有多个，但是只能有原型，不能有具体实现
type d interface {
	pri()
}
type person struct {
	name string
}

func (p *person) pri() {
	fmt.Println(p.name, ":的方法")
}
func hhh(b d) { //多态，传入不同的结构体对象地址，调用不同方法
	b.pri()
}
func main() {
	p1 := person{
		name: "大卫",
	}
	p2 := person{
		name: "小揪揪",
	}
	p1.pri() //大卫
	var a d
	//多态
	//任何实现了这个接口的类型，都可以赋值给这个接口
	//接口内【方法所属于的结构体】可以赋值给该接口
	//即【结构体对象】可以赋值给【实现该结构体方法】的【接口的对象】
	a = &p1
	a.pri() //大卫
	a = &p2
	a.pri()  //小揪揪
	hhh(&p1) //大卫
	hhh(&p2) //小揪揪

}
