/*
append() 和 copy() 函数
如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
下面的代码描述了从拷贝切片的 copy 方法和向切片追加新元素的 append 方法。
*/

package main

import "fmt"

func main() {
	var numbers []int
	printSlice(numbers)
	//len=0 cap=0 slice=[]
	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)
	//len=1 cap=1 slice=[0]
	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)
	//len=2 cap=2 slice=[0 1]
	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)
	//len=5 cap=6 slice=[0 1 2 3 4]
	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
	//len=5 cap=12 slice=[0 1 2 3 4]
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
