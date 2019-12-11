package models

import (
	_ "github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"managerdb/conf"
	"os"
)

var engine *xorm.Engine
var err error

func init() {
	// ---------- xorm start --------------------
	engine, err = xorm.NewEngine("mysql", "root:123456@tcp(localhost:3306)/managerdb?charset=utf8") //managerdb
	if err != nil {
		panic(err)
	}
}
func InitDb() {
	logPath := conf.Global.LogPath.Logdb
	f, err := os.Create(logPath)
	if err != nil {
		panic(err)
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))
	engine.ShowSQL(true)
	engine.ShowExecTime(true)
	engine.SetMaxOpenConns(15)
	engine.SetMaxIdleConns(10)

	pingCount := 1
PingLocal:
	err = engine.Ping()
	if err != nil {
		pingCount++
		goto PingLocal
		if pingCount > 3 {
			panic(err.Error())
		}
	}

	//go func() {
	//	dbs := conf.Global.Ignoredbs // "information_schema,performance_schema,mysql,sys,test"
	//	r, _ := engine.Query("show databases")
	//	for _, b := range r {
	//		for k, a := range b {
	//			if !strings.Contains(dbs, string(a)) {
	//				fmt.Println(k, string(a))
	//			}
	//		}
	//	}
	//}()
	//
	//go func() {
	//	r, _ := engine.Query("show tables")
	//	for _, b := range r {
	//		for k, a := range b {
	//			fmt.Println(k, string(a))
	//		}
	//	}
	//}()

	//cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(),1000)
	//engine.MapCacher(,cacher)
	// -------------------- end --------------------
}

/* beego orm */
//func TableName(name string) string {
//	pre := beego.AppConfig.String("db_dt_prefix")
//	return pre + name
//}
//func GetDBUserName() string {
//	return TableName("db_user")
//}
