package main
import(
	"fmt"
	"net"
	"bufio"
	"log"
	"os"
	"strings"
)

func client(){
	conn,err:=net.Dial("tcp","127.0.0.1:8088")
	if err!=nil {
		log.Fatal(err)
	}

	//客户端可以发送单行数据，然后退出
	reader:=bufio.NewReader(os.Stdin)//os.Stdin表示：标准输入（终端）
	for{
		//从客户端读取一行输入，并准备发送给服务器端
		line,err:=reader.ReadString('\n')
		if err!=nil {
			log.Fatal(err)
		}
		line=strings.Trim(line,"\r\n")
		if line=="exit" {
			fmt.Println("用户退出客户端！")
			break
		}
		//将line发送给服务器端
		conent,err:=conn.Write([]byte(line+"\n"))
		if err!=nil {
			log.Fatal(err)
		}
		fmt.Printf("客户端发送了%d字节的数据到服务器端\n",conent)
	}
}

func main(){
	client()
}