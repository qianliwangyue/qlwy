package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func scan(path string, f os.FileInfo, err error) error {
	fmt.Printf("Scaned: %s\n", path)
	return nil
}
func main() {
	err := os.MkdirAll("fff/sss/gg", 0777) //创建多级目录
	if err != nil {
		fmt.Println(err)
	}

	root := `./fff`

	err = filepath.Walk(root, scan) //err为scan函数返回的的结果
	fmt.Printf("filepath.Walk() returned %v\n", err)

	time.Sleep(time.Second * 5)

	er := os.RemoveAll("fff")
	if er != nil {
		fmt.Println(er)
	}

}
