package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// RSA 数字签名 Demo

// GenerateRsaKey 生成公钥私钥
func GenerateRsaKey(keySize int) {
	// 1. 使用 rsa 包中的 GenerateKey 生成密钥
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		panic(err)
	}
	// 2. 通过 x509 标准对生成的私钥进行序列化
	derText := x509.MarshalPKCS1PrivateKey(privateKey)
	// 3. 将得到的私钥放到 Block 中
	block := &pem.Block{
		Type: "RSA_PRIVATE_KEY",
		Bytes: derText,
	}
	// 4. 执行 pem 编码，并存入磁盘文件
	file,err := os.Create("./数字签名/privateKey.pem")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = pem.Encode(file,block)

	// ==================================== 公钥 ==================================
	// 1. 从私钥中拿到公钥
	publicKey := privateKey.PublicKey
	// 2. 使用 x509 包对公钥进行序列化，此处记住一定要传指针
	derStream, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	// 3. 将得到的公钥放到 Block 中
	block = &pem.Block{
		Type: "RSA_PUBLIC_KEY",
		Bytes: derStream,
	}
	// 4. 执行 pem 编码，并存入磁盘文件
	publicFile, err := os.Create("./数字签名/publicKey.pem")
	if err != nil {
		panic(err)
	}
	err = pem.Encode(publicFile, block)
}

// SignatureRsa RSA 签名
func SignatureRsa(plainText []byte, fileName string) []byte {
	// 1. 打开磁盘文件读取公钥
	file,err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	// 2. 将私钥产品从文件中读出
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	file.Close()
	// 3. 使用 pem 包对使用进行解码
	block, _ := pem.Decode(buf)
	// 4. 使用 x509 将数据解码，得到 pem.Block  结构体变量
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 5. 创建哈希对象
	myHash := sha512.New()
	// 6. 给哈希对象添加数据
	myHash.Write(plainText)
	// 7. 计算哈希值
	hashText := myHash.Sum(nil)
	// 8. 使用 rsa 中的 SignPKCS1v15 函数对散列值签名
	sigText, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA512, hashText)
	if err != nil {
		panic(err)
	}
	return sigText
}

// VerifyRSA RSA 验证签名
func VerifyRSA(plainText,sigText []byte, fileName string) bool {
	// 1. 打开磁盘磁盘中的公钥文件
	file, err := os.Open(fileName)
	if err != nil {panic(err)}
	// 2. 读取文件中的公钥
	fileInfo, err := file.Stat()
	if err != nil {panic(err)}
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	file.Close()
	// 3. 使用 pem 解码，得到 pem.Block 结构体变量
	block, _ := pem.Decode(buf)
	// 4. 使用 x509 对 pem.Block 中的 Bytes 变量中的数进行解析，得到一个接口
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {panic(err)}
	// 5. 进行类型断言，得到公钥结构体
	publicKey := pub.(*rsa.PublicKey)
	// 6. 对原始数据进行哈希运算
	myHash := sha512.New()
	myHash.Write(plainText)
	hashText := myHash.Sum(nil)
	// 7. 进行签名认证
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA512, hashText, sigText)
	if err == nil {
		return true
	}
	return false
}

func main() {
	// 生成密钥对
	GenerateRsaKey(2048)
	src := []byte("圣人之道，悟性自足，向外求理于事务者误也！真正的成功只有一个，按照自己的方式，度过这一生！")
	sigText := SignatureRsa(src, "./数字签名/privateKey.pem")
	bl := VerifyRSA(src, sigText, "./数字签名/publicKey.pem")
	if bl {
		fmt.Println("数字签名验证成功.......!", bl)
	}
}
