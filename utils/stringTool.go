package utils

import (
	"encoding/json"
)

func Byte2Struct(b []byte, inter interface{}) bool {
	isOk := true
	//json数据封装到user对象中
	err := json.Unmarshal(b, &inter)
	if err != nil {
		isOk = false
	}
	return isOk
}
