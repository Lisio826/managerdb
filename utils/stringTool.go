package utils

import (
	"encoding/json"
	"math/rand"
	"time"
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

var str string
func CreateRandStr() string {
	str = "0123456789aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
