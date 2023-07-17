package main

import "fmt"

func main() {
	var dw Animal
	//dw.Name = "狗"
	dw.eat("狗")

}

type Animal struct {
	Name string
}

// （结构体对象 结构体）方法名（参数列表）返回值
func (a Animal) eat(name string) int {
	a.Name = name
	if a.Name == "猫" {
		fmt.Printf("%s :吃鱼\n", a.Name)
	}
	if a.Name == "狗" {
		fmt.Printf("%s:吃屎\n", a.Name)
	}
	return 0
}
