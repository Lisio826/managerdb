package models

import (
	"github.com/astaxie/beego/orm"
)

type DbUser struct {
	Id       int    `json:"-" form:"-" orm:"id"`
	UserName string `json:"userName" form:"userName" orm:"userName"`
	UserPwd  string `json:"userPwd" form:"userPwd" orm:"userPwd"`
	UserStatus   int    `json:"-" form:"userStatus" orm:"userStatus"`
}

// TableName 设置BackendUser表名
func (a *DbUser) TableName() string {
	return GetDBUserName()
}

// BackendUserOneByUserName 根据用户名密码获取单条
func FindDBUserOneByUserName(username, userpwd string) (*DbUser, error) {
	u := DbUser{}
	err := orm.NewOrm().QueryTable(GetDBUserName()).Filter("username", username).Filter("userpwd", userpwd).One(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// BackendUserOne 根据id获取单条
func DBUserOne(id int) (*DbUser, error) {
	o := orm.NewOrm()
	m := DbUser{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
