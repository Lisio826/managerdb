package controllers

import (
	"managedb/enums"
	"managedb/models"
	"managedb/utils"
	"strings"
)

type HomeController struct {
	BaseController
}

func (c *HomeController)Login() {
	username := strings.TrimSpace(c.GetString("userName"))
	userpwd := strings.TrimSpace(c.GetString("userPwd"))
	if len(username) == 0 || len(userpwd) == 0 {
		c.jsonResult(enums.JRCodeFailed, "用户名和密码不正确", "")
	}

	username = utils.DecodeRSA(username)
	userpwd = utils.DecodeRSA(userpwd)

	userpwd = utils.String2md5(userpwd)

	dbuser, err := models.FindDBUserOneByUserName(username,userpwd)
	if err != nil || dbuser == nil {
		c.jsonResult(enums.JRCodeFailed,"用户名或密码错误","")
	}
	if dbuser != nil{
		if dbuser.Status == enums.Disabled{
			c.jsonResult(enums.JRCodeFailed, "用户被禁用，请联系管理员", "")
		}
		//保存用户信息到session
		c.setDBUser2Session(dbuser.Id)
		//获取用户信息
		c.jsonResult(enums.JRCodeSucc, "登录成功", "")
	}


}
