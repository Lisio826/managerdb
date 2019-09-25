package models

import (
	"github.com/astaxie/beego"
	"github.com/go-xorm/xorm"
	"os"
)

var engine *xorm.Engine
var err error
func init() {
	// ---------- xorm start --------------------
	engine, err = xorm.NewEngine("mysql", "root:123456@tcp(localhost:3306)/managerdb?charset=utf8")
	f, err := os.Create("./logs/sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(12)
	engine.SetMaxIdleConns(6)

	//cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(),1000)
	//engine.MapCacher(,cacher)
	// -------------------- end --------------------
}

func TableName(name string) string  {
	pre := beego.AppConfig.String("db_dt_prefix")
	return pre + name
}

func GetDBUserName() string {
	return TableName("db_user")
}
