package service

import (
	"encoding/base64"
	"managerdb/dbmodels"
	"managerdb/enums"
	"managerdb/logger"
	"managerdb/utils"
	"strings"
)

func Login(user *dbmodels.TManageUser) (jsonResult dbmodels.JsonResult) {
	//username := utils.DecodeRSA(user.UserName)
	//userpwd := utils.DecodeRSA(user.UserPwd)
	//私钥
	decodeBytesId, _ := base64.StdEncoding.DecodeString(user.Account)
	code, _ := utils.RsaDecrypt(decodeBytesId) //RSA解密
	usercode := string(code)
	decodeBytesPwd, _ := base64.StdEncoding.DecodeString(user.UserPwd)
	pwd, _ := utils.RsaDecrypt(decodeBytesPwd) //RSA解密
	userpwd := string(pwd)

	userpwd = strings.ToUpper(utils.String2md5(userpwd+enums.PwdSalt))
	logger.Debug(userpwd)
	dbuser, _ := dbmodels.FindDBUserOneByUserName(usercode,userpwd)
	if dbuser != nil{
		if dbuser.UserStatus == enums.UserDisabled{
			jsonResult = dbmodels.JsonResult{enums.JRCodeFailed, "用户被禁用，不可登录", ""}
			return
		}
		jsonResult = dbmodels.JsonResult{enums.JRCodeSucess, "登录成功", ""}
		return
	}
	jsonResult = dbmodels.JsonResult{enums.JRCodeFailed,"用户名或密码错误",""}
	return
}
