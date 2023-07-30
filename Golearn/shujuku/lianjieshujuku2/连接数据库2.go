package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB //定义全局变量

// 定义一个初始化数据库函数
func initSQL() (err error) {
	//连接数据库
	db, err = sql.Open("mysql", "root:19992018470.Zd@tcp(127.0.0.1:3306)/student")
	if err != nil {
		return err
	}
	//尝试连接
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
func main() {
	err := initSQL()
	if err != nil {
		fmt.Printf("init db failed ! ，err:%v\n", err)
	}
	fmt.Println("数据库连接成功！")
}
