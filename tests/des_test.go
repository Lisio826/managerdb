package test

import (
	"crypto/dsa"
	"crypto/rand"
	"fmt"
	"testing"
)

func Test_Des(t *testing.T) {
	/*
		DSA只能做签名 验签 无法做加密
	*/
	var params dsa.Parameters
	//生成参数
	dsa.GenerateParameters(&params, rand.Reader, dsa.L1024N160)
	//生成私钥
	var priv dsa.PrivateKey
	priv.Parameters = params
	dsa.GenerateKey(&priv, rand.Reader)

	//根据私钥生成公钥
	pub := priv.PublicKey

	//利用私钥签名数据
	msg := []byte("hello world")
	r, s, _ := dsa.Sign(rand.Reader, &priv, msg)

	//公钥验签
	b := dsa.Verify(&pub, msg, r, s)
	if b == true {
		fmt.Println("验证成功")
	}
}
