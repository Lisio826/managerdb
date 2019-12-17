package controllers

import (
	"managerdb/dbmodels"
	"managerdb/enums"
	"managerdb/service"
	"managerdb/utils"
)

type HomeController struct {
	BaseController
}
//func (c *HomeController) URLMapping() {
//	c.Mapping("Login", c.Login)
//}

// @router /v1/home/login [*]
func (c *HomeController)Login() {
	//username := strings.TrimSpace(c.GetString("userName"))
	//userpwd := strings.TrimSpace(c.GetString("userPwd"))
	//var user models.DBUser
	//if err := c.ParseForm(&user); err != nil {
	//	c.jsonResult(enums.JRCodeFailed, "用户名或密码不正确", "")
	//}
	//if len(username) == 0 || len(userpwd) == 0 {
	//	c.jsonResult(enums.JRCodeFailed, "用户名或密码不正确", "")
	//}

	var user dbmodels.TManageUser
	data := c.Ctx.Input.RequestBody
	//json数据封装到user对象中
	ok := utils.Byte2Struct(data, &user)
	if !ok && (len(user.Account) == 0 || len(user.UserPwd) == 0) {
		c.jsonResult(enums.JRCodeFailed, "用户名或密码不正确", "")
	}
	jsonResult := service.Login(&user)

	if enums.JRCodeSucess == jsonResult.Code {
		//保存用户信息到session beego的orm用法
		c.setDBUser2Session(&user)
		//c.SetSession("db_user",dbuser)
		////删除指定的session
		//c.DelSession("loginuser")
		////销毁全部的session
		//c.DestroySession()
		//c.Data["json"]=map[string]interface{}{"islogin":islogin};
		//c.ServeJSON();
		/////////////////////////////////

		//c.SetSecureCookie()
		randStr := utils.CreateRandStr()
		mp := make(map[string]string,0)
		mp[enums.Jti] = randStr
		mp[enums.Account] = user.Account
		mp[enums.Key] = randStr
		token := utils.CreateJWT(mp)
		c.Ctx.SetCookie("rand",randStr,"/")
		c.Ctx.SetCookie("token",token,"/")
		//c.SetSession("rand",randStr)
		//c.SetSession("token",token)
	}
	c.jsonResult(jsonResult.Code,jsonResult.Msg,jsonResult.Obj)
}
