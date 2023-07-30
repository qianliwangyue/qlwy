

package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User2 用户信息
type stuinfo2 struct {
	//gorm.Model
	sno   int
	sname  string
	classname string
}


func main() {
	db, err := gorm.Open("mysql", "root:19992018470.Zd@tcp(127.0.0.1:3306)/student?" +
		"charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// 自动迁移
	//db.AutoMigrate(&UserInfo{})

	u1 := &stuinfo2{199901152, "laoshu", "网络21301",}
	//u2 := UserInfo{7,"hunya", "18777777777"}
	// 创建记录
	db.Create(u1)
	//db.Create(&u2)
	// 查询
	//var u = new(UserInfo)
	//db.First(u)
	//fmt.Printf("%#v\n", u)
	//var uu UserInfo
	//db.Find(&uu, "name=?", "laoshu")
	//fmt.Printf("%#v\n", uu)
	//// 更新
	//db.Model(&u).Update("name", "milaoshu")
	//// 删除
	//db.Delete(&u)
}
