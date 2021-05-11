package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"os"
)

// GenerateRsaKey 生成 RSA 的密钥对，并且保存到磁盘文件中
func GenerateRsaKey(keySize int) {
	// 1. 使用 rsa 包中的 GenerateKey 方法生成密钥
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		panic(err)
	}
	// 2. 通过 x509 标准将得到的 rsa 私钥序列化为 ASN.1 的 DER 编码字符串
	derText := x509.MarshalPKCS1PrivateKey(privateKey)
	// 3. 组织一个 pem Block，然后通过 pem 对数据进行编码，通过取地址的方式创建 block
	block := &pem.Block{
		Type: "rsa_private_key", // 这个值写个字符串就行
		//Headers:   // 可选的头项
		Bytes: derText, // 内容解码后的数据，是DER编码的ASN.1结构
	}
	// 4. 执行 pem 编码
	file, err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = pem.Encode(file, block)

	// ======================= 公钥 =======================
	// 1. 从得到的私钥对象中获取公钥信息
	publicKey := privateKey.PublicKey
	// 2. 通过 x509 标准将公钥序列化为 ASN.1 的 DER 编码字符，注此处必须传变量
	derStream, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	// 3. 对序列化后的公钥放到 pem.Block 中
	block = &pem.Block{
		Type: "rsa_public_key",
		//Headers:  // 可选头项
		Bytes: derStream,
	}
	// 4. 进行 pem 编码，并存入到磁盘文件
	publicFile, err := os.Create("public.pem")
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()
	err = pem.Encode(publicFile, block)
}

// RSAEncrypt RSA 公钥加密
func RSAEncrypt(plainText []byte, fileName string) (cipherText []byte) {
	// 1. 读取文件中的公钥
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	// 2. 对读取出来的公钥进行 pem 解码操作
	block, _ := pem.Decode(buf)
	// 3. 使用 x509 进行解码
	publicInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	// 对解码出来的内容进行断言是否为公钥
	if pubKey, ok := publicInterface.(*rsa.PublicKey); ok {
		// 4. 使用公钥加密
		cipherText, err = rsa.EncryptPKCS1v15(rand.Reader, pubKey, plainText)
		if err != nil {
			panic(err)
		}
	}
	return
}

// RSADecrypt RSA 私钥解密
func RSADecrypt(cipherText []byte, fileName string) []byte {
	// 1. 打开文件读取私钥内容
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	// 2. 对读取出来的私钥进的 pem 解密
	block, _ := pem.Decode(buf)
	// 3. 使用 x509 进行解密
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 4. 通过解密出来的私钥对密文进行解密
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		panic(err)
	}
	return plainText
}

// 测试
func main() {
	GenerateRsaKey(1024)
	src := []byte("圣人之道，悟性自足，向外求理于事务者误也！")
	cipherText := RSAEncrypt(src, "public.pem")
	plainText := RSADecrypt(cipherText, "private.pem")
	fmt.Println("======>对公钥解密内容解密出的数据：", string(plainText))

	myHash()
}

// 使用 sha256 对数据进行哈希运算
func myHash() {
	// 方式一：
	// res := sha256.Sum256([]byte("hello, world"))

	// 方式二：
	// 1. 创建一个sha256 的 hash 对象
	myHash := sha256.New()
	// 2. 添加数据
	src := []byte("圣人之道，悟性自足，向外求理于事务者误也！")
	myHash.Write(src)
	myHash.Write(src)
	myHash.Write(src)
	// 3. 计算结果
	res := myHash.Sum(nil)
	// 4. 格式化
	myStr := hex.EncodeToString(res)
	fmt.Println("sha256哈希后的结果：", myStr)
}