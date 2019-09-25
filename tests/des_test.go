package test

import (
	"fmt"
	"github.com/zyx4843/gojson"
	"io/ioutil"
	"managerdb/log"
	"math/rand"
	"net/http"
	"strings"
	"testing"
	"time"
)

func getData(url string) string {
	client := &http.Client{}
	resp, err := client.Get(url)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

type rep struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Results string `json:"results"`
}

// 函　数：生成随机数
// 概　要：
// 参　数：
//      min: 最小值
//      max: 最大值
// 返回值：
//      int64: 生成的随机数
func randInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}
func randInt32(min, max int32) int32 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int31n(max-min) + min
}

func Test_Des(t *testing.T) {

	i := 1
	for i <= 15 {

		log.PprofLogger.Info(fmt.Sprintf(" ------------------------------------ DB %d ------------------------------------ ",i))
		log.MonitorLogger.Info(fmt.Sprintf(" ------------------------------------ DB %d ------------------------------------ ",i))

		pprof := fmt.Sprintf("http://172.22.125.%d:4345/v1/maotai/pprof/?time=%d", i, time.Now().Unix())
		monitor := fmt.Sprintf("http://172.22.125.%d:4345/v1/maotai/monitor?time=%d", i, time.Now().Unix())

		log.PprofLogger.Info(pprof)
		str := getData(pprof)
		str = str[0:strings.LastIndex(str,"Profile Descriptions:")] + "</p></body> </html>"
		log.PprofLogger.Info(str)

		log.PprofLogger.Info(monitor)
		ret := getData(monitor)
		rr := gojson.Json(ret).Get("results").Getdata()
		re1 := rr["NumGoroutine"]
		re2 := rr["DbStatus"]
		re3 := rr["ApiStatus"]
		re4 := rr["MessageChan"]
		re5 := rr["PanicError"]
		//re1 := "NumGoroutine : " + rr.Getdata()["NumGoroutine"]
		log.MonitorLogger.Info("NumGoroutine》》" + fmt.Sprint(re1))
		//mm := rr.Getdata()
		//re2 := mm["DbStatus"]
		log.MonitorLogger.Info("DbStatus》》" + fmt.Sprint(re2))
		//re3 := "ApiStatus : " + rr.Get("ApiStatus").Tostring()
		log.MonitorLogger.Info("ApiStatus》》" + fmt.Sprint(re3))
		//re4 := "MessageChan : " + rr.Get("MessageChan").Tostring()
		log.MonitorLogger.Info("MessageChan》》" + fmt.Sprint(re4))

		log.MonitorLogger.Info("PanicError》》" + fmt.Sprint(re5))
		i = i + 1
		if i == 16 {

			log.PprofLogger.Info(" ------------------------------------ 分割线 ------------------------------------ ")
			log.MonitorLogger.Info(" ------------------------------------ 分割线 ------------------------------------ ")

			time.Sleep(time.Second * 21)
			i = 1
		}
		//log.PprofLogger.Info("测试使用的0 。。。。。。。。。。。。。。。。。。。。。。。")
	}

	//keyword := 山东省济南市槐荫区蓝翔高级中学高一271班团支部"
	//rs := []rune(keyword)
	//keyword = string(rs[0:8])
	//log.Print(keyword)
	//
	//ident := "340122199901142624"
	//month := ident[10:12]
	//year := ident[6:10]
	//fmt.Println(month,year)

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
