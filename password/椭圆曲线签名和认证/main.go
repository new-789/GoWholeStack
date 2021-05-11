package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
)

// GenerateEccKey 生成Ecc密钥对
func GenerateEccKey() {
	/* 1. 使用 ecdsa 生成密钥对,
	参数1:为曲线，可在 elliptic 包中为我们提供好的四个曲线方法，
 	参数2：为 crypto/rand 包中的 Reader
	 */
	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {panic(err)}
	// 2. 使用 x509 进行序列化
	derText, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {panic(err)}
	// 3. 将得到的切片字符放入 pem.BLock 结构体中
	block := &pem.Block{
		Type: "ecdsa_private_key",
		Bytes: derText,
	}
	// 4. 使用 pem 进行编码,并存入文件
	file, err := os.Create("./椭圆曲线签名和认证/eccPrivateKey.pem")
	if err != nil {panic(err)}
	pem.Encode(file, block)
	file.Close()
	// ======================= 公钥 ==========================
	// 1. 从私钥中读取公钥
	publicKey := privateKey.PublicKey
	// 2. 使用 X509 包对公钥进行编码，此处记住一定要传指针类型数据
	derText, err = x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {panic(err)}
	// 3. 将得到的切片字符串放入 pem.Block 中
	block = &pem.Block{
		Type: "ecdsa_publicK_ey",
		Bytes: derText,
	}
	// 4. 使用 pem 进行编码，并存入文件
	pubFile, err := os.Create("./椭圆曲线签名和认证/eccPublicKey.pem")
	if err != nil {panic(err)}
	pem.Encode(pubFile, block)
	pubFile.Close()
}


// EccSignature ecc 私钥签名
func EccSignature(plainText []byte, privName string) (rText, sText []byte) {
	/*
	参数1：需要签名的明文
	参数2：需要打开的存储私钥的文件名
	返回值：签名后并序列化为切片类型的数据
	 */
	// 1. 打开文件读出私钥
	file, err := os.Open(privName)
	if err != nil {panic(err)}
	fileInfo, err := file.Stat()
	if err != nil {panic(err)}
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	file.Close()
	// 2. 使用 pem 对得到的内容私钥内容进行解码
	block, _ := pem.Decode(buf)
	// 3. 使用 x509 对私钥进行反序列化还原操作
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	// 4. 对原始数据进行哈希运算
	myHash := sha256.New()
	myHash.Write(plainText)
	hashText := myHash.Sum(nil)
	// 5. 进行数字签名
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashText)
	if err != nil {panic(err)}
	// 6. 对 r、s 内存中的数据进行格式化 -->byte
	rText, err = r.MarshalText()
	if err != nil {panic(err)}
	sText, err = s.MarshalText()
	if err != nil {panic(err)}
	return
}

// EccVerify ecc 签名认证
func EccVerify(plainText, rText, sText []byte, pubName string) bool {
	/* 参数说明
	参数1：原始数据明文
	参数2、3：使用椭圆曲线私钥签名后得到的切片数据，在做验证前需将其还原为 *big.Int 类型
	参数4：需要打开的存储公钥的文件名
	返回值：验证成功与否 bool 类型数据
	 */
	// 1. 打开文件读取公钥
	file,err := os.Open(pubName)
	if err != nil {panic(err)}
	fileInfo, err := file.Stat()
	if err != nil {panic(err)}
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	file.Close()
	// 2. 使用 pem 对公钥进行解码
	block, _ := pem.Decode(buf)
	// 3. 使用 x509 对数据，并对其进行类型断言得到公钥
	publicInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	publicKey := publicInterface.(*ecdsa.PublicKey)
	if err != nil {panic(err)}
	// 4. 对原始数据进行哈希运算
	myHash := sha256.New()
	myHash.Write(plainText)
	hsaText := myHash.Sum(nil)
	// 5. 使用 ecdsa 进行数据验证，验证之前需现将 sText rText 还原为 *big.Int 类型数据
	var r, s big.Int
	err = r.UnmarshalText(rText)
	err = s.UnmarshalText(sText)
	bl := ecdsa.Verify(publicKey,hsaText,&r, &s)
	return bl
}

func main() {
	// 生成密钥对
	GenerateEccKey()
	// 私钥签名
	src := []byte("圣人之道，悟性自足，向外求理于事务者误也！真正的成功只有一个，按照自己的方式，度过这一生！")
	rText, sText := EccSignature(src, "./椭圆曲线签名和认证/eccPrivateKey.pem")
	// 签名认证
	bl := EccVerify(src, rText, sText, "./椭圆曲线签名和认证/eccPublicKey.pem")
	if bl {
		fmt.Println("椭圆曲线数字签名认证成功", bl)
	}
}
