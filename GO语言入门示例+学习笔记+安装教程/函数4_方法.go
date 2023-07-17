package main

import "fmt"

/* 定义结构体 */
type Circle struct {
	R  float64 //结构体成员1号 R
	PI float64 //成员2号 PI
	S  float64
}

func main() {
	var r1, t Circle //定义r1为成员R的实参
	r1.R = 10.00
	t.PI = 3.14
	fmt.Println("圆的面积 = ", r1.getr2()*t.getPI())
}

//该 method 属于 Circle 类型对象R的方法,即成员对象方法。
func (r Circle) getr2() float64 {
	//r.R 即为 Circle 类型对象中的属性
	return r.R * r.R
}

//该 method 属于 Circle 类型对象PI的方法
func (t Circle) getPI() float64 {
	//t.PI 即为 Circle 类型对象中的属性
	return t.PI
}
