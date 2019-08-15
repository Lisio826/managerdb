package models

import (
	"github.com/astaxie/beego"
)

func init() {

}

func TableName(name string) string  {
	pre := beego.AppConfig.String("db_dt_prefix")
	return pre + name
}

func GetDBUserName() string {
	return TableName("db_user")
}
