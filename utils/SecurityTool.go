package utils

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

var prik = []byte(`-----BEGIN PRIVATE KEY-----
MIICXQIBAAKBgQDn5Yn0vWX1Fr3OwbQWqHgRxG4N6AHKU16Ad4+uy5vw7PSJRce6
sR8cte0HW0KOv7nvl+bBBrs3gpMenUdkmN+HjkQBUlyKVfmFSNvoTpEcdn2vu2UR
jMoRCVEfza/ry9nI6MgsVHGZmOof/t1NofHVoLQki55wN6/bNeOnBRGsXQIDAQAB
AoGAb/K51LKSQ+1EmEmevMl7nWgskP4NzzTMDEyryoB0uaxKqPJM522WTW/uC30c
9njMNEQqm8i6HKQmjcLzsjaywXypjRTEa1Yy37diDbgVpV/xWZk17pFD+HdimYrU
ll82ZkafmGgCNhKHLoPRoUcpovcLR5tSZrb1pIoqj+0BTqkCQQDsJdxfQUjLjRal
qY2JM7B714afrI9/0HNq/+f+7Ia7C1VYr04OS+ELCl1rdv7Y65WIDR5/Z2673TXo
OQ+Q13DTAkEA+2Qvzk9pyHOaGmcu8jJ57qbLoV1FB+OkKV6om7iGtj1f09MHrGtb
I7gaZKlo+q7P7Ql+ClJGIy9tZHWbhPuwDwJBAJm9sHJHg4gZ69Ogxme7wjtuPtQ3
uRkCchIIV1btUG332/Gn+A5wsivI7LcpOpOpFKoFuIRDp6EhTJZKh+rJiEcCQE0E
PpkoPzJIKFgacImG6VAyDYScPH/UQADknSdH+w1t9CPDLUCniz6AMqXQOPdEAzON
iu3CkvZIm20BkunE6gUCQQDOj3i5ow2aCdGjMwk2OupeyCkTypxLUPUYZ2VZNeqz
7EOoUlW+8Iu6i2QO5mOUu5O13Y4vJHRQRERnjnJlSRc/
-----END PRIVATE KEY-----`)

func DecodeRSA(str string) string {
	//pem解码
	block, _ := pem.Decode(prik)
	//x509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err!=nil{
		panic(err)
	}
	plaintext, e := rsa.DecryptOAEP(md5.New(), rand.Reader, privateKey, []byte(str), nil)
	if e != nil {
		fmt.Println(e)
	}
	return string(plaintext)
}
