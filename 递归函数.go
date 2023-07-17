/*

递归，就是在运行的过程中调用自己。

语法格式如下：

func recursion() {
   recursion() 函数调用自身
}

func main() {
   recursion()
}
注意：Go 语言支持递归。但我们在使用递归时，开发者需要设置退出条件，否则递归将陷入无限循环中。

递归函数对于解决数学上的问题是非常有用的，就像计算阶乘，生成斐波那契数列等。
*/

package main

import "fmt"

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	var i int = 15
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))

	fmt.Println("=============================================")
	var a int
	for a = 0; a < 10; a++ {
		fmt.Printf("%d\t", fibonacci(a))
	}
}

//f0 f1 f2 f3	f4	f5	f6	f7	f8	返回函数前两之和
//0 1	2	3	5	8	13	21	34
