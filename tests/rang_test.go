package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var r *rand.Rand

func init()  {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}
// RandString 生成随机字符串
func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func TestRand(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandString(10))
	}
}

func Test_rand( t *testing.T)  {
	str := "0123456789aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	fmt.Println(string(result))
}

func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min) + min
	return randNum
}

func A()  {
	fmt.Println(rand.Intn(200) - 100)
}
