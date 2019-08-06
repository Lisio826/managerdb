package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"managedb/controllers"
	_ "managedb/routers"
)

func main() {

	beego.BConfig.Log.AccessLogs = true
	beego.BConfig.Log.FileLineNum = true
	beego.BConfig.Log.Outputs = map[string]string{"console": ""}

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
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		//公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		//如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))
	//路由设置
	ns := beego.NewNamespace("/v1",
		//  用于跨域请求
		beego.NSRouter("*", &controllers.BaseController{}, "OPTIONS:Options"), )
	beego.AddNamespace(ns)

	beego.Run()
}
