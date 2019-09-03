package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql"
	_ "managerdb/log"
	"managerdb/models"
	_ "managerdb/models"
	_ "managerdb/routers"
	"time"
)

func init() {
	////跨域设置
	//var FilterGateWay = func(ctx *context.Context) {
	//	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	//	//允许访问源
	//	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")
	//	//允许post访问
	//	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,ContentType,Authorization,accept,accept-encoding, authorization, content-type") //header的类型
	//	ctx.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
	//	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	//}
	//beego.InsertFilter("*", beego.BeforeRouter, FilterGateWay)

}

func main() {

	//ok := logs.ConfigLog()
	//	//if !ok {
	//	//	return
	//	//}

	//orm.RegisterModel(new(models.DBUser))
	//orm.RegisterDriver("mysql",orm.DRMySQL)
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(localhost:3306)/managerdb?charset=utf8", maxIdle, maxConn)
	orm.DefaultTimeLoc = time.UTC

	orm.RegisterModel(new(models.DbUser))
	orm.Debug = true

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "sessionID"

	// TODO 支持跨域访问
	//InsertFilter是提供一个过滤函数
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//允许访问所有源
		AllowAllOrigins: true,
		//可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		//其中Options跨域复杂请求预检
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		//指的是允许的Header的种类
		AllowHeaders: []string{"_c", "Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"}, //"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"
		//公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"}, //"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"
		//如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
		//AllowOrigins:     []string{"*"},
	}))

	beego.Run()
}
