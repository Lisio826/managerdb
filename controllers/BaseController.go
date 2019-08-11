package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"managerdb/enums"
	"managerdb/models"
)

type BaseController struct {
	beego.Controller
	controllerName string        //当前控制名称
	actionName     string        //当前action名称
	curUser        models.DBUser //当前用户信息
}

func (c *BaseController) Options() {
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}

func (c *BaseController) Prepare() {
	//c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	//附值
	c.controllerName, c.actionName = c.GetControllerAndAction()

	fmt.Print(c)

	//log.Print(strings.TrimSpace(c.GetString("userName")))
	//log.Print(strings.TrimSpace(c.GetString("userPwd")))
	//var u models.DBUser
	//log.Print(c.Ctx.Input.RequestBody)
	//log.Print(json.Unmarshal(c.Ctx.Input.RequestBody,&u))
	//从Session里获取数据 设置用户信息
	//c.adapterUserInfo()
}

//// checkLogin判断用户是否登录，未登录则跳转至登录页面
//// 一定要在BaseController.Prepare()后执行
//func (c *BaseController) checkLogin() {
//	if c.curUser.Id == 0 {
//		//登录成功后返回的址为当前
//		returnURL := c.Ctx.Request.URL.Path
//
//		if c.Ctx.Input.IsPost() {
//			//returnURL := c.Ctx.Input.Refer()
//			c.jsonResult(enums.JRCode302, "请登录", returnURL)
//		}
//		c.StopRun()
//	}
//}

func (c *BaseController) jsonResult(code enums.JsonResultCode, msg string, obj interface{}) {
	r := &models.JsonResult{code, msg, obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

////SetBackendUser2Session 获取用户信息（包括资源UrlFor）保存至Session
//func (c *BaseController) setDBUser2Session(userId int) error {
//	m, err := models.DBUserOne(userId)
//	if err != nil {
//		return err
//	}
//	////获取这个用户能获取到的所有资源列表
//	//resourceList := models.ResourceTreeGridByUserId(userId, 1000)
//	//for _, item := range resourceList {
//	//	m.ResourceUrlForList = append(m.ResourceUrlForList, strings.TrimSpace(item.UrlFor))
//	//}
//	c.SetSession("db_user", *m)
//	return nil
//}

////从session里取用户信息
//func (c *BaseController) adapterUserInfo() {
//	s := c.GetSession("db_user")
//	if s != nil {
//		c.curUser = s.(models.DBUser)
//		c.Data["db_user"] = s
//	}
//}
