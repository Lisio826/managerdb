package test

import (
	"fmt"
	"github.com/astaxie/beego"
	"gopkg.in/redis.v5"
	_ "managerdb/routers"
	"path/filepath"
	"runtime"
	"testing"
	//. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)


}


// TestBeego is a sample to run an endpoint test
func TestBeego(t *testing.T) {
	ssname := "获取%s数据失败"
	ssname = fmt.Sprintf(ssname,"福建")
	fmt.Println(ssname)


	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)


	//r, _ := http.NewRequest("GET", "/", nil)
	//w := httptest.NewRecorder()
	//beego.BeeApp.Handlers.ServeHTTP(w, r)
	//
	//beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())
	//
	//Convey("Subject: Test Station Endpoint\n", t, func() {
	//        Convey("Status Code Should Be 200", func() {
	//                So(w.Code, ShouldEqual, 200)
	//        })
	//        Convey("The Result Should Not Be Empty", func() {
	//                So(w.Body.Len(), ShouldBeGreaterThan, 0)
	//        })
	//})
}

