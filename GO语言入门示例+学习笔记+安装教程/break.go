package main

import "fmt"

func main() {

	// 不使用标记
	fmt.Println("---- break ----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break //不使用标记，只是跳出内循环，外循环将继续
		}
	}

	// 使用标记
	fmt.Println("===========================================================")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break re //直接跳出标记的外循环
		}
	}
}

/*
---- break ----
i: 1
i2: 11
i: 2
i2: 11
i: 3
i2: 11
===========================================================
i: 1
i2: 11
*/
