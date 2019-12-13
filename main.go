package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"managerdb/conf"
	"managerdb/logger"
	"managerdb/models"
	_ "managerdb/routers"
)

/**
 * 使用 beego 作为服务框架，MVC模式
 * 数据驱动用 xorm（可批量查询） ：https://blog.csdn.net/kenkao/article/details/91429282
 * 使用 cap 用来记录日志
 * 使用 ants 用来作为线程池
 * 使用 redis 存储部分数据
 * 验证使用 session + jwt
 */
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

	conf.Global.GetConf()
	logger.InitLog()
	models.InitDb()
}

func main() {

	//ok := logs.ConfigLog()
	//	//if !ok {
	//	//	return
	//	//}

	//// --------- beego 自带的org start --------------------
	////orm.RegisterModel(new(models.DBUser))
	////orm.RegisterDriver("mysql",orm.DRMySQL)
	//maxIdle := 30
	//maxConn := 30
	//orm.RegisterDataBase("default", "mysql", "root:123456@tcp(localhost:3306)/managerdb?charset=utf8", maxIdle, maxConn)
	//orm.DefaultTimeLoc = time.UTC
	//
	//orm.RegisterModel(new(models.DbUser))
	//orm.Debug = true
	//// --------------------- end -------------------------

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
