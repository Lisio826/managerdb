package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"managerdb/dbmodels"
	"managerdb/enums"
	"managerdb/logger"
	"strings"
)
type BaseController struct {
	beego.Controller
	controllerName string        //当前控制名称
	actionName     string        //当前action名称
	curUser        dbmodels.TManageUser //当前用户信息
}
//var l *zap.SugaredLogger
//func InitLog() {
//	fmt.Println("================================")
//	l = logger.MainLogger
//}
func (c *BaseController) Options() {
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}

func (c *BaseController) Prepare() {
	//c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	//附值
	c.controllerName, c.actionName = c.GetControllerAndAction()
	logger.Debug(c.controllerName,c.actionName)

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
func (c *BaseController) checkLogin() {
	if c.curUser.Id == 0 {
		//登录成功后返回的址为当前
		returnURL := c.Ctx.Request.URL.Path

		if c.Ctx.Input.IsPost() {
			//returnURL := c.Ctx.Input.Refer()
			c.jsonResult(enums.JRCode302, "请登录", returnURL)
		}
		c.StopRun()
	}
}
var FilterUser = func(ctx *context.Context){
	_, ok := ctx.Input.Session("uid").(string)
	ok2 := strings.Contains(ctx.Request.RequestURI, "/login")
	if !ok && !ok2{
		ctx.Redirect(302, "/login/index")
	}
}
// ParseToken parse JWT token in http header.
func (c *BaseController) ParseToken() (t *jwt.Token, e error) {
	authString := c.Ctx.Input.Header("Authorization")
	logger.Debug("AuthString:", authString)

	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		logger.Error("AuthString invalid:", authString)
		return nil, errors.New("无效的验证码")
	}
	tokenString := kv[1]

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("mykey"), nil
	})
	if err != nil {
		logger.Error("Parse token:", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// That's not even a token
				return nil, errInputData
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return nil, errExpired
			} else {
				// Couldn't handle this token
				return nil, errInputData
			}
		} else {
			// Couldn't handle this token
			return nil, errInputData
		}
	}
	if !token.Valid {
		logger.Error("Token invalid:", tokenString)
		return nil, errInputData
	}
	////////////
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		logger.Error(errPermission)
		return
	}
	var user string = claims["username"].(string)
	/////////////////////////
	return token, nil
}

func (c *BaseController) jsonResult(code enums.JsonResultCode, msg string, obj interface{}) {
	r := &dbmodels.JsonResult{code, msg, obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

////SetBackendUser2Session 获取用户信息（包括资源UrlFor）保存至Session
func (c *BaseController) setDBUser2Session(user *dbmodels.TManageUser) error {
	//m, err := models.DBUserOne(userId)
	//if err != nil {
	//	return err
	//}
	////获取这个用户能获取到的所有资源列表
	//resourceList := models.ResourceTreeGridByUserId(userId, 1000)
	//for _, item := range resourceList {
	//	m.ResourceUrlForList = append(m.ResourceUrlForList, strings.TrimSpace(item.UrlFor))
	//}
	c.SetSession("db_user", user)

	return nil
}

////从session里取用户信息
//func (c *BaseController) adapterUserInfo() {
//	s := c.GetSession("db_user")
//	if s != nil {
//		c.curUser = s.(models.DBUser)
//		c.Data["db_user"] = s
//	}
//}
