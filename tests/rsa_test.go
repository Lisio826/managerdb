package test

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/wenzhenxi/gorsa"
	"log"
	"os"
	"testing"
)

func init() {

	//a := gorsa.RSASecurity{}
	//a.SetPrivateKey("manage_db_pri")
	//a.SetPublicKey("manage_db_pub")
	//
	//fmt.Println(a)
	//
	//
	//fmt.Println("----------------------------------------")

}
var Pubkey = `-----BEGIN 公钥-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAk+89V7vpOj1rG6bTAKYM
56qmFLwNCBVDJ3MltVVtxVUUByqc5b6u909MmmrLBqS//PWC6zc3wZzU1+ayh8xb
UAEZuA3EjlPHIaFIVIz04RaW10+1xnby/RQE23tDqsv9a2jv/axjE/27b62nzvCW
eItu1kNQ3MGdcuqKjke+LKhQ7nWPRCOd/ffVqSuRvG0YfUEkOz/6UpsPr6vrI331
hWRB4DlYy8qFUmDsyvvExe4NjZWblXCqkEXRRAhi2SQRCl3teGuIHtDUxCskRIDi
aMD+Qt2Yp+Vvbz6hUiqIWSIH1BoHJer/JOq2/O6X3cmuppU4AdVNgy8Bq236iXvr
MQIDAQAB
-----END 公钥-----
`

var Pirvatekey = `-----BEGIN 私钥-----
MIIEpAIBAAKCAQEAk+89V7vpOj1rG6bTAKYM56qmFLwNCBVDJ3MltVVtxVUUByqc
5b6u909MmmrLBqS//PWC6zc3wZzU1+ayh8xbUAEZuA3EjlPHIaFIVIz04RaW10+1
xnby/RQE23tDqsv9a2jv/axjE/27b62nzvCWeItu1kNQ3MGdcuqKjke+LKhQ7nWP
RCOd/ffVqSuRvG0YfUEkOz/6UpsPr6vrI331hWRB4DlYy8qFUmDsyvvExe4NjZWb
lXCqkEXRRAhi2SQRCl3teGuIHtDUxCskRIDiaMD+Qt2Yp+Vvbz6hUiqIWSIH1BoH
Jer/JOq2/O6X3cmuppU4AdVNgy8Bq236iXvrMQIDAQABAoIBAQCCbxZvHMfvCeg+
YUD5+W63dMcq0QPMdLLZPbWpxMEclH8sMm5UQ2SRueGY5UBNg0WkC/R64BzRIS6p
jkcrZQu95rp+heUgeM3C4SmdIwtmyzwEa8uiSY7Fhbkiq/Rly6aN5eB0kmJpZfa1
6S9kTszdTFNVp9TMUAo7IIE6IheT1x0WcX7aOWVqp9MDXBHV5T0Tvt8vFrPTldFg
IuK45t3tr83tDcx53uC8cL5Ui8leWQjPh4BgdhJ3/MGTDWg+LW2vlAb4x+aLcDJM
CH6Rcb1b8hs9iLTDkdVw9KirYQH5mbACXZyDEaqj1I2KamJIU2qDuTnKxNoc96HY
2XMuSndhAoGBAMPwJuPuZqioJfNyS99x++ZTcVVwGRAbEvTvh6jPSGA0k3cYKgWR
NnssMkHBzZa0p3/NmSwWc7LiL8whEFUDAp2ntvfPVJ19Xvm71gNUyCQ/hojqIAXy
tsNT1gBUTCMtFZmAkUsjqdM/hUnJMM9zH+w4lt5QM2y/YkCThoI65BVbAoGBAMFI
GsIbnJDNhVap7HfWcYmGOlWgEEEchG6Uq6Lbai9T8c7xMSFc6DQiNMmQUAlgDaMV
b6izPK4KGQaXMFt5h7hekZgkbxCKBd9xsLM72bWhM/nd/HkZdHQqrNAPFhY6/S8C
IjRnRfdhsjBIA8K73yiUCsQlHAauGfPzdHET8ktjAoGAQdxeZi1DapuirhMUN9Zr
kr8nkE1uz0AafiRpmC+cp2Hk05pWvapTAtIXTo0jWu38g3QLcYtWdqGa6WWPxNOP
NIkkcmXJjmqO2yjtRg9gevazdSAlhXpRPpTWkSPEt+o2oXNa40PomK54UhYDhyeu
akuXQsD4mCw4jXZJN0suUZMCgYAgzpBcKjulCH19fFI69RdIdJQqPIUFyEViT7Hi
bsPTTLham+3u78oqLzQukmRDcx5ddCIDzIicMfKVf8whertivAqSfHytnf/pMW8A
vUPy5G3iF5/nHj76CNRUbHsfQtv+wqnzoyPpHZgVQeQBhcoXJSm+qV3cdGjLU6OM
HgqeaQKBgQCnmL5SX7GSAeB0rSNugPp2GezAQj0H4OCc8kNrHK8RUvXIU9B2zKA2
z/QUKFb1gIGcKxYr+LqQ25/+TGvINjuf6P3fVkHL0U8jOG0IqpPJXO3Vl9B8ewWL
cFQVB/nQfmaMa4ChK0QEUe+Mqi++MwgYbRHx1lIOXEfUJO+PXrMekw==
-----END 私钥-----
`
// 公钥加密私钥解密
func applyPubEPriD() error {
	pubenctypt, err := gorsa.PublicEncrypt(`hello world`,Pubkey)
	if err != nil {
		return err
	}
	pridecrypt, err := gorsa.PriKeyDecrypt(pubenctypt,Pirvatekey)
	if err != nil {
		return err
	}
	fmt.Println(string(pridecrypt))
	if string(pridecrypt) != `hello world` {
		return errors.New(`解密失败`)
	}
	return nil
}

// 公钥解密私钥加密
func applyPriEPubD() error {
	prienctypt, err := gorsa.PriKeyEncrypt(`hello world`,Pirvatekey)
	if err != nil {
		return err
	}
	pubdecrypt, err := gorsa.PublicDecrypt(prienctypt,Pubkey)
	if err != nil {
		return err
	}
	if string(pubdecrypt) != `hello world` {
		return errors.New(`解密失败`)
	}
	return nil
}

func test3()  {
	// 公钥加密私钥解密
	if err := applyPubEPriD(); err != nil {
		log.Println(err)
	}
	// 公钥解密私钥加密
	if err := applyPriEPubD(); err != nil {
		log.Println(err)
	}
}
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
//rsa加密解密 签名验签
func test2() {
	//生成私钥
	priv, e := rsa.GenerateKey(rand.Reader, 1024)
	if e != nil {
		fmt.Println(e)
	}
	X509PrivateKey,_ := x509.MarshalPKIXPublicKey(priv)
	//privv := pem.EncodeToMemory(&pem.Block{Type: "RSA Private Key",Bytes:X509PrivateKey})

	// 保存到文件
	privateFile, err := os.Create("tests/private_2.pem")
	if err!=nil{
		panic(err)
	}
	defer privateFile.Close()
	//构建一个pem.Block结构体对象
	privateBlock:= pem.Block{Type: "PRIVATE KEY",Bytes:X509PrivateKey}
	//将数据保存到文件
	pem.Encode(privateFile,&privateBlock)

	////pem解码
	//block, _ := pem.Decode(prik)
	////x509解码
	//privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	//if err!=nil{
	//	panic(err)
	//}

	//根据私钥产生公钥
	pub := &priv.PublicKey

	pubby,_ := x509.MarshalPKIXPublicKey(pub)
	// 这个生成的密钥在 jsencrypt 中 使用 失败
	//pubby,_ := x509.MarshalPKIXPublicKey(pub)
	//保存到文件
	publicFile, err := os.Create("tests/public_2.pem")
	if err!=nil{
		panic(err)
	}
	defer publicFile.Close()
	//构建一个pem.Block结构体对象
	publicBlock:= pem.Block{Type: "PUBLIC KEY",Bytes:pubby}
	//将数据保存到文件
	pem.Encode(publicFile,&publicBlock)

	//明文
	plaintext := []byte("Hello world")
	//加密生成密文
	fmt.Printf("%q\n加密:\n", plaintext)
	ciphertext, e := rsa.EncryptOAEP(md5.New(), rand.Reader, pub, plaintext, nil)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("\t%x\n", ciphertext)

	//解密得到明文
	fmt.Printf("解密:\n")
	plaintext, e = rsa.DecryptOAEP(md5.New(), rand.Reader, priv, ciphertext, nil)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(plaintext))
	fmt.Printf("%q", plaintext)

	//消息先进行Hash处理
	h := md5.New()
	h.Write(plaintext)
	hashed := h.Sum(nil)
	fmt.Printf("\n%q MD5 Hashed:\n\t%x\n", plaintext, hashed)

	//签名
	opts := &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto, Hash: crypto.MD5}
	sig, e := rsa.SignPSS(rand.Reader, priv, crypto.MD5, hashed, opts)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("签名:\n\t%x\n", sig)

	//认证
	fmt.Printf("验证结果:")
	if e := rsa.VerifyPSS(pub, crypto.MD5, hashed, sig, opts); e != nil {
		fmt.Println("失败:", e)
	} else {
		fmt.Println("成功.")
	}
}

func TestRsa(t *testing.T)  {

	//test3()

	test2()

	//test1()
}

func test1()  {
	//生成密钥对，保存到文件
	GenerateRSAKey(2048)
	message:=[]byte("hello world")
	//加密
	cipherText:= RsaEncrypt(message,"tests/public.pem")
	fmt.Println("加密后为：",string(cipherText))
	//解密
	plainText := RsaDecrypt(cipherText, "tests/private.pem")
	fmt.Println("解密后为：",string(plainText))
}

//生成RSA私钥和公钥，保存到文件中
func GenerateRSAKey(bits int){
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err!=nil{
		panic(err)
	}
	//保存私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//使用pem格式对x509输出的内容进行编码
	//创建文件保存私钥
	privateFile, err := os.Create("tests/private.pem")
	if err!=nil{
		panic(err)
	}
	defer privateFile.Close()
	//构建一个pem.Block结构体对象
	privateBlock:= pem.Block{Type: "RSA PRIVATE KEY",Bytes:X509PrivateKey}
	//将数据保存到文件
	pem.Encode(privateFile,&privateBlock)

	//保存公钥
	//获取公钥的数据
	publicKey:=privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey,err:=x509.MarshalPKIXPublicKey(&publicKey)
	if err!=nil{
		panic(err)
	}
	//pem格式编码
	//创建用于保存公钥的文件
	publicFile, err := os.Create("tests/public.pem")
	if err!=nil{
		panic(err)
	}
	defer publicFile.Close()
	//创建一个pem.Block结构体对象
	publicBlock:= pem.Block{Type: "RSA PUBLIC KEY",Bytes:X509PublicKey}
	//保存到文件
	pem.Encode(publicFile,&publicBlock)
}

//RSA加密
func RsaEncrypt(plainText []byte,path string)[]byte{
	//打开文件
	file,err:=os.Open(path)
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	//读取文件的内容
	info, _ := file.Stat()
	buf:=make([]byte,info.Size())
	file.Read(buf)

	//base64
	//encodeString := base64.StdEncoding.EncodeToString(buf)
	//fmt.Println(encodeString)

	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err!=nil{
		panic(err)
	}
	//类型断言
	publicKey:=publicKeyInterface.(*rsa.PublicKey)



	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err!=nil{
		panic(err)
	}
	//返回密文
	return cipherText
}

//RSA解密
func RsaDecrypt(cipherText []byte,path string) []byte{
	//打开文件
	file,err:=os.Open(path)
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	//获取文件内容
	info, _ := file.Stat()
	buf:=make([]byte,info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err!=nil{
		panic(err)
	}
	//对密文进行解密
	plainText,_:=rsa.DecryptPKCS1v15(rand.Reader,privateKey,cipherText)
	//返回明文
	return plainText
}


//将字符串加密成 md5
func String2md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}
// 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func Test(t *testing.T)  {
	fmt.Println(String2md5("abc123_Ab123"))
	//fmt.Println(MD5("abc"))
}
