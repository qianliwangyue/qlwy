package main
import "fmt"

func main() {
	fmt.Println("hello go")
	var hh uint
	hh = 15
	fmt.Println(hh)
	Animal()
	//var re uint16
	re := max(15, 68)
	fmt.Printf("%d\n", re) //自定义格式化输出
	fmt.Println(re)        //换行输出
	fmt.Print(re)          //默认输出
}
func Animal() {
	fmt.Println("狗吃你")
}
func max(n1, n2 uint16) uint16 {
	var result uint16
	if n1 > n2 {
		result = n1
	} else {
		result = n2
	}
	return result
}
