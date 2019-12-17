package routers

import (
	"github.com/astaxie/beego"
	"managerdb/controllers"
	"strings"
)

var FilterUser = func(ctx *context.Context){
	_, ok := ctx.Input.Session("uid").(string)
	ok2 := strings.Contains(ctx.Request.RequestURI, "/login")
	if !ok && !ok2{
		ctx.Redirect(302, "/login/index")
	}
}

func init() {

	beego.InsertFilter("/",beego.BeforeRouter, FilterUser)

	beego.Router("/", &controllers.MainController{})

	//路由设置
	ns := beego.NewNamespace("/v1",
		//  用于跨域请求
		beego.NSRouter("*", &controllers.BaseController{}, "OPTIONS:Options"), )
	beego.AddNamespace(ns)

    //beego.Router("/v1/home/login",&controllers.HomeController{},"*:Login")
	beego.Include(&controllers.HomeController{})
}
