package test

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)
//https://blog.csdn.net/wangshubo1989/article/details/74529333
//https://www.cnblogs.com/lrj567/p/6209872.html
var str = "dfasfasfasfasdgfafasf"
//iss: 签发者
//sub: 面向的用户
//aud: 接收方
//exp: 过期时间
//nbf: 生效时间
//iat: 签发时间
//jti: 唯一身份标识
func TestCreateT(tt *testing.T)  {
	t := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["iat"] = time.Now().Unix()
	claims["nbf"] = time.Now().Unix()
	//claims["exp"] = time.Now().Add(time.Hour * time.Duration(2)).Unix()
	claims["exp"] = time.Now().Add(time.Second * time.Duration(2)).Unix()

	claims["admin"] = 0
	claims["name"] = "username"
	t.Claims = claims

	ts,err := t.SignedString([]byte(str))
	if err == nil {
		fmt.Println(ts)
	}

	fmt.Println("=======================================")
	time.Sleep(time.Second * 1)

	to,err := jwt.Parse(ts, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(str),nil
	})

	if to.Valid {
		fmt.Println(to)
		cl,ok := to.Claims.(jwt.MapClaims)
		if ok {
			fmt.Println(cl["name"].(string))
		}
	}else{
		fmt.Println("无效 token ")
	}

}

func parseT(ts string) (error) {  //ControllerError
	to,err := jwt.Parse(ts, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(str),nil
	})
	if err != nil {
		beego.Error("Parse token:", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// That's not even a token
				return errors.New("nihaoo")	//errInputData
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return errors.New("nihaoo")	//errExpired
			} else {
				// Couldn't handle this token
				return errors.New("nihaoo")	//errInputData
			}
		} else {
			// Couldn't handle this token
			return errors.New("nihaoo")	//errInputData
		}
	}

	if to.Valid {
		fmt.Println(to)
	}else{
		fmt.Println("无效 token ")
	}

	return nil

}
