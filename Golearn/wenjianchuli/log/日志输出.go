package main

import (
	"log"
	"os"
)

func main() {
	fileName := "目录/New.log"
	logFile, err := os.Create(fileName) //创建文件
	if err != nil {
		log.Fatalln("open file error")
	}
	defer logFile.Close()
	debugLog := log.New(logFile, "[Info]", log.Llongfile) //文件名，标题名，日志文件的路径和运行信息
	debugLog.Println("Info Level Message")                //输出的内容
	debugLog.SetPrefix("[Debug]")                         //标题名
	debugLog.Println("Debug Level Message")               //输出的内容
}

//log.Print\Println\Printf  	处理一般日志exit(0)
//log.Panic   					处理意外日志exit(2)
//log.Fatal\Fatalln\Fatalf		处理致命日志exit(1)
