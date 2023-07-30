package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	limitchan = make(chan struct{}, 10)
)

func limitiumiddleware() gin.HandlerFunc { //限流中间件
	return func(ctx *gin.Context) {
		limitchan <- struct{}{}
		ctx.Next()
		<-limitchan
	}
}
func timemiddleware() gin.HandlerFunc { //计时中间件，显示运行所用时间
	return func(ctx *gin.Context) {
		begin := time.Now()
		ctx.Next() //继续执行后面的所有函数，即：hander1和hander2
		fmt.Printf("use time: %d ms\n", time.Since(begin).Milliseconds())
	}
}

func hander1(ctx *gin.Context) {
	time.Sleep(100 * time.Millisecond)
	ctx.String(http.StatusOK, "高性能golang。\n")
}
func hander2(ctx *gin.Context) {
	time.Sleep(100 * time.Millisecond)
	ctx.String(http.StatusOK, "高性能golang。\n")
}

func main() {
	engine := gin.Default()
	//engine.Use(timemiddleware(), limitiumiddleware()) //设置全局中间件,以后下面每个路由都自动配备此中间件
	//注意Use（中间件1，中间件2,中间件3...）
	//将中间件1放到路由最前面，将中间件2放到路由最前面，将中间件3放到路由最前面，下面路由执行顺序3,2,1
	engine.GET("/1", timemiddleware(), limitiumiddleware(), hander1, hander2)
	//GET“/1”路由上单独配置的中间件，timemiddleware(), limitiumiddleware()
	//访问“/”页面显示
	//高性能golang。
	//高性能golang。
	//输出显示：use time: 201 ms//hander1和hander2各有100ms延迟
	engine.Run(":5678")
}
