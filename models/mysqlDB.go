package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func RegisterDB() {
	orm.RegisterModel(new(User), new(Reward), new(Suggestion))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
}
