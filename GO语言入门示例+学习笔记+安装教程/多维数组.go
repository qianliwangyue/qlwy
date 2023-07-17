package main

import "fmt"

func main() {
	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)

	// Step 3: 显示两行数据
	fmt.Println("Row 1")
	fmt.Println(values[0])
	fmt.Println("Row 2")
	fmt.Println(values[1])

	// Step 4: 访问第一个元素
	fmt.Println("第一个元素为：")
	fmt.Println(values[0][0])
	//row1和row2组成二维数组values[2][3]
	fmt.Println("=================================================================")
	/* 数组 - 3 行 4 列*/
	a := [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}} /* 第三行索引为 2 */
	fmt.Println(a[2][3]) //11
	fmt.Println(a[1][2]) //6
	fmt.Println(a)       //输出二维数组a

	fmt.Println("=================================================================")

	/* 数组 - 5 行 2 列*/
	var b = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int

	/* 输出数组元素 */
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, b[i][j])
		}
	}

	fmt.Println("=================================================================")

	// 创建空的二维数组
	animals := [][]string{}

	// 创建三一维数组，各数组长度不同
	row3 := []string{"fish", "shark", "eel"}
	row4 := []string{"bird"}
	row5 := []string{"lizard", "salamander"}

	// 使用 append() 函数将一维数组添加到二维数组中
	animals = append(animals, row3)
	animals = append(animals, row4)
	animals = append(animals, row5)

	// 循环输出
	for i := range animals {
		fmt.Printf("Row: %v\n", i)
		fmt.Println(animals[i])
	}
	fmt.Println("=================================================================")
	/* 数组长度为 5 */
	var balance = [5]int{1000, 2, 3, 17, 50}
	var avg float32

	/* 数组作为参数传递给函数 */
	avg = getAverage(balance, 5)

	/* 输出返回的平均值 */
	fmt.Printf("平均值为: %f ", avg)

	fmt.Println("=================================================================")
	x := 1.69
	y := 1.7
	z := x * y // 结果应该是2.873
	//浮点数计算输出有一定的偏差，你也可以转整型来设置精度。
	fmt.Println(z) // 输出的是2.8729999999999998

	x = 1690                          // 表示1.69                       // 表示1.69
	y = 1700                          // 表示1.70
	z = x * y                         // 结果应该是2873000表示 2.873
	fmt.Println(z)                    // 内部编码
	fmt.Println(float64(z) / 1000000) // 显示
}
func getAverage(arr [5]int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}
