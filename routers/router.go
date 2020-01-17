package routers

import (
	"github.com/astaxie/beego"
	"managerdb/controllers"

)

func init() {

	beego.InsertFilter("/",beego.BeforeRouter, controllers.FilterUser)


	//beego.Router("/", &controllers.MainController{})

	//路由设置
	ns := beego.NewNamespace("/v1",
		//  用于跨域请求
		beego.NSRouter("*", &controllers.BaseController{}, "OPTIONS:Options"), )
	beego.AddNamespace(ns)

    //beego.Router("/v1/home/login",&controllers.HomeController{},"*:Login")
	beego.Include(&controllers.HomeController{})
}
