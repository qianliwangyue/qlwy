package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 表的设计
type User struct {
	Id      int
	Name    string `orm:"unique"`
	Pwd     string
	Article []*Article `orm:"rel(m2m)"`
}

// 文章结构体
type Article struct {
	Id       int       `orm:"pk;auto"`                 //主键，自增
	ArtiName string    `orm:"size(20)"`                //长度20
	Atime    time.Time `orm:"auto_now;type(datetime)"` //每次model保存都会自动更新时间
	// `orm:"auto_now_add;type(date)"` 第一次保存时才设置时间

	Acount   int `orm:"default(0);null"` //默认为0,允许为空
	Acontent string
	Aimg     string //文件路径
	Atype    string //文件路径

	ArticleType *ArticleType `orm:"rel(fk)"`       //RelForeignKey
	User        []*User      `orm:"reverse(many)"` //对应的反向关系 RelReverseMany
}

// 类型表
type ArticleType struct {
	Id      int
	Tname   string
	Article []*Article `orm:"reverse(many)"`
}

func init() {
	// 设置数据库基本信息
	orm.RegisterDataBase("default", "mysql", "root:19992018470.Zd@tcp(127.0.0.1:3306)/beego1?charset=utf8")
	// 映射model数据
	orm.RegisterModel(new(User), new(Article), new(ArticleType))
	// 生成表
	orm.RunSyncdb("default", false, true)
}
