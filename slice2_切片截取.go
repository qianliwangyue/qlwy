// 可以通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]，实例如下：
package main

import "fmt"

func main() {
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)
	//len=9 cap=9 slice=[0 1 2 3 4 5 6 7 8]
	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)
	//numbers == [0 1 2 3 4 5 6 7 8]
	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])
	//numbers[1:4] == [1 2 3]
	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])
	//numbers[:3] == [0 1 2]
	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])
	//numbers[4:] == [4 5 6 7 8]
	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)
	//len=0 cap=5 slice=[]
	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)
	//len=2 cap=9 slice=[0 1]
	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)
	//len=3 cap=7 slice=[2 3 4]
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
