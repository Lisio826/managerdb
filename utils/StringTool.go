package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

//将字符串加密成 md5
func String2md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}

func Byte2Struct(b []byte, inter interface{}) bool {
	isOk := true
	//json数据封装到user对象中
	err := json.Unmarshal(b, &inter)
	if err != nil{
		isOk = false
	}
	return isOk
}
