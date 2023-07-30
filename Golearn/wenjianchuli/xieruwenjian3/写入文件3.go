package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./wenjianchuli/duquwenjian1/格式化写入.txt")
	if err != nil {
		fmt.Println("创建失败:", err)
		return
	}
	defer file.Close()
	for i := 0; i < 5; i++ {
		aaa := fmt.Sprintf("%s:%d\n", "Hello Go!", i) //Sprintf格式化
		file.WriteString(aaa)                         //string信息
		file.Write([]byte("i love go\n\n"))           //byte类型
	}

}
