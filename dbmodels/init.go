package dbmodels

import (
	"fmt"
	_ "github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gopkg.in/redis.v5"
	"managerdb/conf"
	"os"
	"strings"
	"time"
)

var engine *xorm.Engine
var err error

func InitRedisServer() {
	iRedis := conf.Global.RedisServer
	client := redis.NewClient(&redis.Options{
		Addr:         iRedis.Addr,
		Password:     iRedis.Password,        // no password set
		DB:           iRedis.DB, // use default DB
		MaxRetries:   iRedis.MaxRetries,
		DialTimeout:  time.Second * time.Duration(iRedis.DialTimeout),
		ReadTimeout:  time.Second * time.Duration(iRedis.ReadTimeout),
		WriteTimeout: time.Second * time.Duration(iRedis.WriteTimeout),
		PoolSize:     iRedis.PoolSize,
		PoolTimeout:  time.Second * time.Duration(iRedis.PoolTimeout),
		IdleTimeout:  time.Second * time.Duration(iRedis.IdleTimeout),
	})
	pong, err := client.Ping().Result()
	if err != nil || strings.ToUpper(pong) != "PONG" {
		panic(err)
	}
}
func InitDb() {
	// ---------- xorm start --------------------
	iMysql := conf.Global.MysqlServer
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",iMysql.Username,iMysql.Password,iMysql.Ip,iMysql.Database)
	engine, err = xorm.NewEngine("mysql", dataSourceName) //managerdb
	if err != nil {
		panic(err)
	}
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
