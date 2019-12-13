package enums

type JsonResultCode int
const (
	JRCodeSucc JsonResultCode = iota
	JRCodeFailed
	JRCode302 = 302 //跳转至地址
	JRCode401 = 401 //未授权访问

	PwdSalt = "_Ab123"

	UserEnabled = true
	UserDisabled = false

	Jti = "jti"
	Account = "account"
	Key = "key"
	Admin = "admin"
)

const (
	Deleted = iota - 1
	Disabled
	Enabled
)
