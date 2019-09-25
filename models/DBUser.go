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
	Id          int       `db:"id" json:"id"`
	FullName    string    `db:"real_name" json:"real_name"`
	Surname     string    `db:"surname" json:"surname"`
	Name        string    `db:"name" json:"name"`
	UserCode    string    `db:"user_code" json:"user_code"`
	Identity    string    `db:"identity" json:"identity"`
	UserStatus  bool      `db:"user_status" json:"user_status"`
	UserPwd     string    `db:"user_pwd" json:"user_pwd"`
	Mobile      string    `db:"mobile" json:"mobile"`
	Email       string    `db:"email" json:"email"`
	RoleId      int       `db:"role_id" json:"role_id"`
	Avatar      string    `db:"avatar" json:"avatar"`
	AddTime     time.Time `db:"add_time" json:"add_time"`
	OperateUser string    `db:"operate_user" json:"operate_user"`
	UpdateTime  time.Time `db:"update_time" json:"update_time"`
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
