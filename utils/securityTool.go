package utils

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"managerdb/logger"
	"time"
)

var privateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQCXrYqCy3fYvwoSNBZ9q0xc6EVxlGTQWr9hm1hrpCCid1/C+kX2
sZlT1YJo+IZ47KaY+tN+sEXujTugWT7BJzz44bv2NUELZCEoNCsHpXsGJh0FzGcJ
Sscg7W1ZBQmRcdHX/zhUNcOodnQIkfXw7ebqAHi2B8mc3VwH8IqqwRLUuwIDAQAB
AoGAARG6Hj3s9Fw0Z/hyCCTC/wysIKbZJxbcroV31BpuGHca97bn81tBJ3Qj3TWT
TRrt+boa75OyioH4X6MBbD8okN3O1Yhi4u6NhReNVs3WyqYPNjf6wqbEn0QD/p1u
cVKdsUX5GWS8rk6r+h4RDoCPfQu8jJEZQdocwDgPs+6disECQQC+Ncd36gJn32Ad
Dwm8K6BXmex6u1IpUT5Bo6OACT47u3MRAgndvzHbwCEm48Kdnm3uBhnDoaMvxYXC
NNqQyreRAkEAzCPqdi5omnubKtGLHIM2yigSAnG87uJ5CUnp7JagM075njakM2ef
96LX4QAiVdRW0mUkxep/A7IH04EJo9QZiwJAFXiwqaKq3triWKVn7evmz7y6ayhW
H//hzTsRq9RfSBLi7FqUjomUsT/A/kjboTtPIX1izyih6TSc4s1gltOrYQJAXolr
dRt82b3QV/3+XE/vuaLg5Dzj8AT1ZHXUXPpMajZypvBLDns4u6ewLcSwIetZ4clx
RC6ab42szif93N7GjQJAOlzCV631Kp9yh/WvTULJkBFlSlBmn+fA36YxCWQ4mHbu
MnsbowOguxYs97ZlP3h4u17P4a+1/ou4Tj1jyqQFgQ==
-----END RSA PRIVATE KEY-----`)

func RsaDecrypt(cipherText []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey) //将密钥解析成私钥实例
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, cipherText) //RSA算法解密
}

var prik = []byte(`-----BEGIN PRIVATE KEY-----
MIICXQIBAAKBgQDn5Yn0vWX1Fr3OwbQWqHgRxG4N6AHKU16Ad4+uy5vw7PSJRce6
sR8cte0HW0KOv7nvl+bBBrs3gpMenUdkmN+HjkQBUlyKVfmFSNvoTpEcdn2vu2UR
jMoRCVEfza/ry9nI6MgsVHGZmOof/t1NofHVoLQki55wN6/bNeOnBRGsXQIDAQAB
AoGAb/K51LKSQ+1EmEmevMl7nWgskP4NzzTMDEyryoB0uaxKqPJM522WTW/uC30c
9njMNEQqm8i6HKQmjcLzsjaywXypjRTEa1Yy37diDbgVpV/xWZk17pFD+HdimYrU
ll82ZkafmGgCNhKHLoPRoUcpovcLR5tSZrb1pIoqj+0BTqkCQQDsJdxfQUjLjRal
qY2JM7B714afrI9/0HNq/+f+7Ia7C1VYr04OS+ELCl1rdv7Y65WIDR5/Z2673TXo
OQ+Q13DTAkEA+2Qvzk9pyHOaGmcu8jJ57qbLoV1FB+OkKV6om7iGtj1f09MHrGtb
I7gaZKlo+q7P7Ql+ClJGIy9tZHWbhPuwDwJBAJm9sHJHg4gZ69Ogxme7wjtuPtQ3
uRkCchIIV1btUG332/Gn+A5wsivI7LcpOpOpFKoFuIRDp6EhTJZKh+rJiEcCQE0E
PpkoPzJIKFgacImG6VAyDYScPH/UQADknSdH+w1t9CPDLUCniz6AMqXQOPdEAzON
iu3CkvZIm20BkunE6gUCQQDOj3i5ow2aCdGjMwk2OupeyCkTypxLUPUYZ2VZNeqz
7EOoUlW+8Iu6i2QO5mOUu5O13Y4vJHRQRERnjnJlSRc/
-----END PRIVATE KEY-----`)

func DecodeRSA(str string) string {
	//pem解码
	block, _ := pem.Decode(prik)
	//x509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	plaintext, e := rsa.DecryptOAEP(md5.New(), rand.Reader, privateKey, []byte(str), nil)
	if e != nil {
		fmt.Println(e)
	}
	return string(plaintext)
}

//将字符串加密成 md5
func String2md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}

//func MD5(text string) string {
//	ctx := md5.New()
//	ctx.Write([]byte(text))
//	return hex.EncodeToString(ctx.Sum(nil))
//}


func CreateJWT(mp map[string]string) string {
	t := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["iat"] = time.Now().Unix()
	claims["nbf"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(2)).Unix()
	//claims["exp"] = time.Now().Add(time.Second * time.Duration(2)).Unix()
	if len(mp) > 0 {
		for k,v := range mp{
			claims[k] = v
		}
	}
	//claims["jti"] = str
	//claims["admin"] = 0
	//claims.["key"] =
	//claims["account"] = "username"
	t.Claims = claims
	ts,err := t.SignedString([]byte(str))
	if err != nil {
		return ""
	}
	return ts
	//fmt.Println("=======================================")
	//time.Sleep(time.Second * 1)
}

func ValidJWT(str string) (jwt.MapClaims,error) {
	t := jwt.New(jwt.SigningMethodHS256)
	ts,err := t.SignedString([]byte(str))
	if err != nil {
		logger.Error(err)
		return nil,errors.New("无效 token, 解析失败")
	}
	to,err := jwt.Parse(ts, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(str),nil
	})
	if to.Valid {
		cl,ok := to.Claims.(jwt.MapClaims)
		if ok {
			//fmt.Println(cl["name"].(string))
			return cl,nil
		}
	}
	return nil,errors.New("无效 token，验证失败")
}

func CreateJWT1(mp map[string]string) (tokenString string) {
	// 带权限创建令牌
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 480).Unix() //20天有效期，过期需要重新登录获取token
	if len(mp) > 0 {
		for k,v := range mp{
			claims[k] = v
		}
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用自定义字符串加密 and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte("mykey"))
	if err != nil {
		beego.Error("jwt.SignedString:", err)
	}
	return
	//fmt.Println("=======================================")
	//time.Sleep(time.Second * 1)
}
