package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// 定义结构体对应XML的元素
type EmailConfig struct {
	XMLName        xml.Name       `xml:"config"`
	SmtpServer     string         `xml:"smtpServer"`
	SmtpPort       int            `xml:"smtpPort"`
	Sender         string         `xml:"sender"`
	SenderPassword string         `xml:"senderPassword"`
	Receivers      EmailReceivers `xml:"receivers"`
}

// 定义结构体对应XML的元素
type EmailReceivers struct {
	Flag string   `xml:"flag,attr"`
	User []string `xml:"user"`
}

func main() {
	file, err := os.Open("目录/email_config.xml") //打开XML文件
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file) //读取文件到内存data
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := EmailConfig{}            //创建空的结构体对象
	err = xml.Unmarshal(data, &v) //将内存data的内容赋值给空结构体对象v
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
	fmt.Println("SmtpServer is : ", v.SmtpServer) //通过访问结构体对象成员，来访问XML文件中的元素
	fmt.Println("SmtpPort is : ", v.SmtpPort)
	fmt.Println("Sender is : ", v.Sender)
	fmt.Println("SenderPasswd is : ", v.SenderPassword)
	fmt.Println("Receivers.Flag is : ", v.Receivers.Flag)
	for i, element := range v.Receivers.User {
		fmt.Println(i, element)
	}
}
