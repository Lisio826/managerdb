package routers

import (
	"github.com/astaxie/beego"
	"managedb/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/v1/home/login",&controllers.HomeController{},"get:Login")

	////初始化 namespace
	//ns :=
	//	beego.NewNamespace("/v1",
	//		beego.NSCond(func(ctx *context.Context) bool {
	//			if ctx.Input.Domain() == "api.beego.me" {
	//				return true
	//			}
	//			return false
	//		}),
	//		beego.NSNamespace("/home",
	//			beego.NSGet("/index", func(ctx *context.Context) {
	//				ctx.Output.Body([]byte("shopinfo"))
	//			}),
	//		),
	//	)
	////注册 namespace
	//beego.AddNamespace(ns)

}
