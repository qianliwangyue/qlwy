package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 45, 4}
	b := a[1:3] //b为切片，是数组的引用（截取数组a第2,3位）
	fmt.Println(b)
	fmt.Println(len(b)) //2
	fmt.Println(cap(b)) //容量：4
	b[1] = 6            //切片的第二位，对应数组a的第3位
	fmt.Println(a)
	fmt.Println("=================================")
	sle := make([]int, 5, 6) //定义初始数组长度为5,初始分配内存容量6
	sle[2] = 10
	sle[4] = 5
	sle[3] = 84
	sle = append(sle, 9) //长度为5的数组，曾加第6个数
	sle = append(sle, 7) //长度为5的数组，曾加第7个数
	//初始分配内存容量6,此时已经曾加了第7个数，内存不够后，自动重新分配内存为原来的两倍，cap(sle)=12
	sle = append(sle, 6, 8, 5)
	fmt.Println(sle)
	fmt.Println(cap(sle))

}

//切片可以截取数组，切片也可以截取切片
//切片可以改变数组，但不能用来改变string，string是一个整体
