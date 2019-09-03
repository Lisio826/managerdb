package logs
//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/astaxie/beego"
//	config2 "github.com/astaxie/beego/config"
//	"github.com/astaxie/beego/logs"
//)
//
//type logstruct struct {
//	filename            string `json:"filename"`
//	maxlines            int    `json:"maxlines"`
//	maxsize             int    `json:"maxsize"`
//	daily               bool   `json:"daily"`
//	maxdays             int    `json:"maxdays"`
//	rotate              bool   `json:"rotate"`
//	level               string `json:"level"`
//	perm                int    `json:"perm"`
//	enableFuncCallDepth bool   `json:"-"`
//	logFuncCallDepth    int    `json:"-"`
//}
//
//func ConfigLog() bool {
//	beego.BConfig.Log.AccessLogs = true
//	beego.BConfig.Log.FileLineNum = true
//
//	logconf, err := config2.NewConfig("ini", "conf/log.conf")
//	if err != nil {
//		fmt.Println(err.Error())
//		return false
//	}
//
//	maxlines, _ := logconf.Int("maxlines")
//	maxsize, _ := logconf.Int("maxsize")
//	daily, _ := logconf.Bool("daily")
//	maxdays, _ := logconf.Int("maxdays")
//	rotate, _ := logconf.Bool("rotate")
//	perm, _ := logconf.Int("perm")
//	enableFuncCallDepth, _ := logconf.Bool("enableFuncCallDepth")
//	logFuncCallDepth, _ := logconf.Int("logFuncCallDepth")
//	config := logstruct{
//		logconf.String("filename"),
//		maxlines,
//		maxsize,
//		daily,
//		maxdays,
//		rotate,
//		logconf.String("level"),
//		perm,
//		enableFuncCallDepth,
//		logFuncCallDepth,
//	}
//
//	//config := make(map[string]interface{})
//	configstr, err := json.Marshal(config)
//	if err != nil {
//		fmt.Println("init Logger failed, marshal err: ", err)
//		return false
//	}
//	logs.EnableFuncCallDepth(config.enableFuncCallDepth)
//	logs.Async()
//
//	logs.SetLogger(logs.AdapterFile, string(configstr))
//	logs.SetLogger(logs.AdapterConsole)
//	return true
//}
