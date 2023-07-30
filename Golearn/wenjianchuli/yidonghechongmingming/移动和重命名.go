package main

import (
	
	"fmt"
	"os"
)
//创建，憨.txt，并移动到 ，"目录"下,且改名为,憨憨阿皮.txt
func main(){
	file,err:=os.Create("憨.txt")
	if err!=nil{
		fmt.Println("创建失败！：",err)
	}
	file.Close()


	err=os.Mkdir("目录",0777)
	chulierr(err)
	err=os.Rename("憨.txt","目录/憨憨阿皮.txt")
	if err!=nil{
		fmt.Println(err)
		return
	}
}
func chulierr(e error){
	if e!=nil {
		fmt.Println(e)
	}
}