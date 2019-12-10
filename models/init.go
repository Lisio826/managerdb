package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-xorm/xorm"
	"os"
	"strings"
)

var engine *xorm.Engine
var err error
func init() {
	// ---------- xorm start --------------------
	engine, err = xorm.NewEngine("mysql", "root:123456@tcp(localhost:3306)/managerdb?charset=utf8") //managerdb
	f, err := os.Create("./logs/sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(12)
	engine.SetMaxIdleConns(6)

	go func() {
		dbs := "information_schema,performance_schema,mysql,sys"
		r,_ := engine.Query("show databases")
		for _,b := range r{
			for k,a := range b{
				if !strings.Contains(dbs,string(a)) {
					fmt.Println(k, string(a))
				}
			}
		}
	}()

	go func() {
		r,_ := engine.Query("show tables")
		for _,b := range r{
			for k,a := range b{
				fmt.Println(k, string(a))
			}
		}}()

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
