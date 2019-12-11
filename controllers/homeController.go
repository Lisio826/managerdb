package controllers

import (
	"encoding/base64"
	"managerdb/enums"
	"managerdb/logger"
	"managerdb/models"
	"managerdb/utils"
)

type HomeController struct {
	BaseController
}
func (c *HomeController) URLMapping() {
	c.Mapping("Login", c.Login)
}

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

	var user models.TManageUser
	data := c.Ctx.Input.RequestBody
	//json数据封装到user对象中
	ok := utils.Byte2Struct(data, &user)
	if !ok && (len(user.Account) == 0 || len(user.UserPwd) == 0) {
		c.jsonResult(enums.JRCodeFailed, "用户名或密码不正确", "")
	}

	//username := utils.DecodeRSA(user.UserName)
	//userpwd := utils.DecodeRSA(user.UserPwd)
	//私钥
	decodeBytesId, _ := base64.StdEncoding.DecodeString(user.Account)
	id, err := utils.RsaDecrypt(decodeBytesId) //RSA解密
	username := string(id)
	decodeBytesPwd, _ := base64.StdEncoding.DecodeString(user.UserPwd)
	pwd, err := utils.RsaDecrypt(decodeBytesPwd) //RSA解密
	userpwd := string(pwd)

	userpwd = utils.String2md5(userpwd+enums.PwdSalt)
	logger.MainLogger.Debug(userpwd)
	dbuser, err := models.FindDBUserOneByUserName(username,userpwd)
	if err != nil || dbuser == nil {
		c.jsonResult(enums.JRCodeFailed,"用户名或密码错误","")
	}
	if dbuser != nil{
		if dbuser.UserStatus == enums.UserDisabled{
			c.jsonResult(enums.JRCodeFailed, "用户被禁用，不可登录", "")
		}
		//保存用户信息到session beego的orm用法
		//c.setDBUser2Session(dbuser.Id)
		c.SetSession("db_user",dbuser)
		////删除指定的session
		//c.DelSession("loginuser")
		////销毁全部的session
		//c.DestroySession()
		//c.Data["json"]=map[string]interface{}{"islogin":islogin};
		//c.ServeJSON();
		/////////////////////////////////

		//获取用户信息
		c.jsonResult(enums.JRCodeSucc, "登录成功", "")
	}


}
