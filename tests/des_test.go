package test

import (
	"fmt"
	"maotai/log"
	"testing"
)

func Test_Des(t *testing.T) {
	keyword := "山东省济南市槐荫区蓝翔高级中学高一271班团支部"
	rs := []rune(keyword)
	keyword = string(rs[0:8])
	log.Print(keyword)

	ident := "340122199901142624"
	month := ident[10:12]
	year := ident[6:10]
	fmt.Println(month,year)

	///*
	//	DSA只能做签名 验签 无法做加密
	//*/
	//var params dsa.Parameters
	////生成参数
	//dsa.GenerateParameters(&params, rand.Reader, dsa.L1024N160)
	////生成私钥
	//var priv dsa.PrivateKey
	//priv.Parameters = params
	//dsa.GenerateKey(&priv, rand.Reader)
	//
	////根据私钥生成公钥
	//pub := priv.PublicKey
	//
	////利用私钥签名数据
	//msg := []byte("hello world")
	//r, s, _ := dsa.Sign(rand.Reader, &priv, msg)
	//
	////公钥验签
	//b := dsa.Verify(&pub, msg, r, s)
	//if b == true {
	//	fmt.Println("验证成功")
	//}
}
