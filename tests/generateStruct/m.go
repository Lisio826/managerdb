package main

import "time"

type My_db_user struct {
	Id         int    `db:"id" json:"id"`
	UserName   string `db:"user_name" json:"user_name"`
	UserPwd    string `db:"user_pwd" json:"user_pwd"`
	UserStatus int    `db:"user_status" json:"user_status"`
}

func (*My_db_user) TableName() string {
	return "my_db_user"
}

type T_manage_permission struct {
	Id               int       `db:"id" json:"id"`
	PermissionName   string    `db:"permission_name" json:"permission_name"`
	PermissionCode   string    `db:"permission_code" json:"permission_code"`
	PermissionType   string    `db:"permission_type" json:"permission_type"`
	PermissionStatus string    `db:"permission_status" json:"permission_status"`
	AddTime          time.Time `db:"add_time" json:"add_time"`
	OperateUser      string    `db:"operate_user" json:"operate_user"`
	UpdataTime       time.Time `db:"updata_time" json:"updata_time"`
}

func (*T_manage_permission) TableName() string {
	return "t_manage_permission"
}

type T_manage_resource struct {
	Id             int       `db:"id" json:"id"`
	ResourceName   string    `db:"resource_name" json:"resource_name"`
	ParentId       int       `db:"parent_id" json:"parent_id"`
	ResourceCode   string    `db:"resource_code" json:"resource_code"`
	ResourceType   string    `db:"resource_type" json:"resource_type"`
	ResourceStatus string    `db:"resource_status" json:"resource_status"`
	AddTime        time.Time `db:"add_time" json:"add_time"`
	OperateUser    string    `db:"operate_user" json:"operate_user"`
	UpdataTime     time.Time `db:"updata_time" json:"updata_time"`
}

func (*T_manage_resource) TableName() string {
	return "t_manage_resource"
}

type T_manage_role struct {
	Id          int       `db:"id" json:"id"`
	RoleName    string    `db:"role_name" json:"role_name"`
	RoleCode    string    `db:"role_code" json:"role_code"`
	RoleStatus  string    `db:"role_status" json:"role_status"`
	AddTime     time.Time `db:"add_time" json:"add_time"`
	OperateUser string    `db:"operate_user" json:"operate_user"`
	UpdataTime  time.Time `db:"updata_time" json:"updata_time"`
}

func (*T_manage_role) TableName() string {
	return "t_manage_role"
}

type T_manage_role_resource_scope struct {
	Id           int       `db:"id" json:"id"`
	RoleId       int       `db:"role_id" json:"role_id"`
	ResouceId    int       `db:"resouce_id" json:"resouce_id"`
	PermissionId int       `db:"permission_id" json:"permission_id"`
	Status       string    `db:"status" json:"status"`
	OperateUser  string    `db:"operate_user" json:"operate_user"`
	AddTime      time.Time `db:"add_time" json:"add_time"`
	IdDelete     string    `db:"id_delete" json:"id_delete"`
}

func (*T_manage_role_resource_scope) TableName() string {
	return "t_manage_role_resource_scope"
}

type T_manage_user struct {
	Id          int       `db:"id" json:"id"`
	RealName    string    `db:"real_name" json:"real_name"`
	SurName     string    `db:"sur_name" json:"sur_name"`
	Name        string    `db:"name" json:"name"`
	UserCode    string    `db:"user_code" json:"user_code"`
	Identity    string    `db:"identity" json:"identity"`
	Mobile      string    `db:"mobile" json:"mobile"`
	Email       string    `db:"email" json:"email"`
	RoleId      int       `db:"role_id" json:"role_id"`
	Avatar      string    `db:"avatar" json:"avatar"`
	AddTime     time.Time `db:"add_time" json:"add_time"`
	OperateUser string    `db:"operate_user" json:"operate_user"`
	UpdateTime  time.Time `db:"update_time" json:"update_time"`
}

func (*T_manage_user) TableName() string {
	return "t_manage_user"
}

type T_manage_user_resouce_scope struct {
	Id           int       `db:"id" json:"id"`
	UserId       int       `db:"user_id" json:"user_id"`
	ResouceId    int       `db:"resouce_id" json:"resouce_id"`
	PermissionId int       `db:"permission_id" json:"permission_id"`
	Status       string    `db:"status" json:"status"`
	OperateUser  string    `db:"operate_user" json:"operate_user"`
	AddTime      time.Time `db:"add_time" json:"add_time"`
	IsDelete     string    `db:"is_delete" json:"is_delete"`
}

func (*T_manage_user_resouce_scope) TableName() string {
	return "t_manage_user_resouce_scope"
}
