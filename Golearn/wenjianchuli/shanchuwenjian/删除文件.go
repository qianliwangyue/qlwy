package main
import(
	"fmt"
	"os"
	"time"
)
func main(){
	err:=os.MkdirAll("目录/目录1",0777)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println("create dir:目录1")
	//remove1
	fp,err:=os.Create("./目录/目录1/remove1.txt")
	if err!=nil {
		fmt.Println("remove1.txt create failed:",err)
	}
	fp.Close()
	fmt.Println("create file remove1.txt")
	//remove2
	fp,err=os.Create("./目录/目录1/remove2.txt")
	if err!=nil {
		fmt.Println("remove2.txt create failed:",err)
	}
	fp.Close()
	fmt.Println("create file remove2.txt")
	//remove3
	fp,err=os.Create("./目录/目录1/remove3.txt")
	if err!=nil {
		fmt.Println("remove3.txt create failed:",err)
	}
	fp.Close()
	fmt.Println("create file remove3.txt")

	time.Sleep(time.Second*10)


	//删除文件
	err=os.Remove("./目录/目录1/remove1.txt")
	if err!=nil {
		fmt.Printf("remove1.txt删除失败:%v\n",err)
	}
	fmt.Println("remove1.txt删除成功！")
	err=os.RemoveAll("./目录/目录1")
	if err!=nil {
		fmt.Println("目录1删除失败:",err)
	}
	fmt.Println("目录1删除成功！")

}