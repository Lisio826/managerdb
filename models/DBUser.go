package models

import (
	"github.com/astaxie/beego/orm"
)

type DBUser struct {
	Id int

	Status int
}

// BackendUserOneByUserName 根据用户名密码获取单条
func FindDBUserOneByUserName(username, userpwd string) (*DBUser, error) {
	u := DBUser{}
	err := orm.NewOrm().QueryTable(GetDBUserName()).Filter("username", username).Filter("userpwd", userpwd).One(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// BackendUserOne 根据id获取单条
func DBUserOne(id int) (*DBUser, error) {
	o := orm.NewOrm()
	m := DBUser{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
