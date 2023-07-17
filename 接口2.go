package main

import "fmt"

//定义了一个 Shape 接口
type Shape interface {
	area() float64 //它定义了一个方法 area()
}

//定义了结构体 Rectangle
type Rectangle struct {
	width  float64
	height float64
}

//area方法返回一个 float64 类型的面积值
func (r Rectangle) area() float64 {
	return r.width * r.height
}

//定义了结构体Circle
type Circle struct {
	radius float64
}

//area方法返回一个 float64 类型的面积值
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	var s Shape

	s = Rectangle{width: 10, height: 5}
	fmt.Printf("矩形面积: %f\n", s.area())

	s = Circle{radius: 3}
	fmt.Printf("圆形面积: %f\n", s.area())
}

/*
在 main() 函数中，我们首先定义了一个 Shape 类型的变量 s，
然后分别将 Rectangle 和 Circle 类型的实例赋值给它，
并通过 area() 方法计算它们的面积并打印出来，输出结果如下：
矩形面积: 50.000000
圆形面积: 28.260000
需要注意的是，接口类型变量可以存储任何实现了该接口的类型的值。
在示例中，我们将 Rectangle 和 Circle 类型的实例都赋值给了 Shape 类型的变量 s，
并通过 area() 方法调用它们的面积计算方法。
*/
