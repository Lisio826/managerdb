package models

import (
	"errors"
	"time"
)

type DbUser struct {
	Id         int    `json:"-" form:"-" orm:"id"`
	UserName   string `json:"userName" form:"userName" orm:"userName"`
	UserPwd    string `json:"userPwd" form:"userPwd" orm:"userPwd"`
	UserStatus int    `json:"-" form:"userStatus" orm:"userStatus"`
}
type TManageUser struct {
	Id          int       `json:"id"`
	FullName    string    `json:"fullName"`
	Surname     string    `json:"surname"`
	Name        string    `json:"name"`
	UserCode    string    `json:"userCode"`
	Identity    string    `json:"identity"`
	UserStatus  bool      `json:"userStatus"`
	UserPwd     string    `json:"userPwd"`
	Mobile      string    `json:"mobile"`
	Email       string    `json:"email"`
	RoleId      int       `json:"roleId"`
	Avatar      string    `json:"avatar"`
	AddTime     time.Time `json:"addTime"`
	OperateUser string    `json:"operateUser"`
	UpdateTime  time.Time `json:"updateTime"`
}

// TableName 设置BackendUser表名
func (a *DbUser) TableName() string {
	return GetDBUserName()
}

// BackendUserOneByUserName 根据用户名密码获取单条
//func FindDBUserOneByUserName(username, userpwd string) (*DbUser, error) {
//	u := DbUser{}
//	err := orm.NewOrm().QueryTable(GetDBUserName()).Filter("username", username).Filter("userpwd", userpwd).One(&u)
//	if err != nil {
//		return nil, err
//	}
//	return &u, nil
//}
func FindDBUserOneByUserName(usercode, userpwd string) (*TManageUser, error) {
	tManageUser := &TManageUser{UserCode: usercode, UserPwd: userpwd}
	has, err := engine.Get(tManageUser)
	if err != nil || !has {
		return nil, errors.New("查无此用户")
	}
	return tManageUser, nil
}

//// BackendUserOne 根据id获取单条
//func DBUserOne(id int) (*DbUser, error) {
//	o := orm.NewOrm()
//	m := DbUser{Id: id}
//	err := o.Read(&m)
//	if err != nil {
//		return nil, err
//	}
//	return &m, nil
//}
