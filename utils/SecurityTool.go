package utils

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

var prik = []byte(`-----BEGIN RSA Private Key-----
MIICXgIBAAKBgQDAwj7wtpKcnjsGUMhPz6tN+86pbtycUpvNqofAShadSnvj7Zte
bD0ZwILUs/KRTHpXCbfuZQuWt0XKr+NhgAHSYa6X31pY9xluH/MT6Sk+N5jpe9kn
X0BlNp7RG75ll5EUR1nDnt8i+ZAjWfu2Nto2zwMAnfuU3ZNvwwedmXMYJQIDAQAB
AoGAJL6imuZiymJLZCfQF47B3ArNWXUdrtr8glcoq5oz92X0ef2YaYZ+m0McbYmk
CtimVaMvoc03SVLEhh/DW5m4PtqfMZSjuMEfVkn0VoAZPU7v/iWoHRxY77O6nIE2
4vg8wWU+WCa9PcKfsiCqLlxRt3lHRTonBJhzLiWgJRx5+50CQQDu8j79vp6bbH7m
pdUjkwEYSasMi6/rH1Yq6LcbZ/OUPlCuICjUF6JgRKCEeIzJsl6bg07HewGxfSAS
ir5Y+1GvAkEAzoQdS6U/mMy6k0RXs3UOAj8YSW3T0dsujEn2mHqWq3Dk4HAvlDkT
i8LhBqUDYksn1MWfKVghQCHlsORByFhMawJBANkv5WOczOu23U3mc+om53rdExWP
LnKHhDnyVuUZQoR/c7Qh1Rqa9OON3V0recSnVWkDHCsjOnHKqslPxR56KU8CQQCv
EVZYWzTHjr6XuzbxnciLZPtsvBr16u3R5Z2Tc1Co82JDVTcwWxZTw5fJbzeoKvgQ
Kpubi+dcG6BAza9qbFgTAkEAw4CsLlliGt7lNueWw6aPVer+VM5mpGFkszFMlYv3
Dd+8qqCc/H1Ic32C2xh8pkNmSj8x/dWLN3pdWYV/ikJrsg==
-----END RSA Private Key-----`)

var pubk = []byte(`-----BEGIN RSA Public Key-----
MIGJAoGBAMDCPvC2kpyeOwZQyE/Pq037zqlu3JxSm82qh8BKFp1Ke+Ptm15sPRnA
gtSz8pFMelcJt+5lC5a3Rcqv42GAAdJhrpffWlj3GW4f8xPpKT43mOl72SdfQGU2
ntEbvmWXkRRHWcOe3yL5kCNZ+7Y22jbPAwCd+5Tdk2/DB52ZcxglAgMBAAE=
-----END RSA Public Key-----`)

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
