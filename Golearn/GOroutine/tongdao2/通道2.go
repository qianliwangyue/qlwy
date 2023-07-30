package main
import(
	"fmt"
	"time"
)

func Echo(s string){
	for i := 0; i < 3; i++ {
		time.Sleep(2*time.Second)
		fmt.Println(s)
	}
}
func main(){
	go Echo("golangman")
	Echo("sigma")
}

