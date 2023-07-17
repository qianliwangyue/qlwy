package main

import "fmt"

const (
	i = 1 << iota //iota初始为0，左移0位
	j = 3 << iota //左移1位：3二进制为0011左移后0110化为二进制：6
	k             //0110左移2位 ：1100    12
	l = 5
)

func main() {
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", 1+l)

}
