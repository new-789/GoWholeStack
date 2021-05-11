package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

// 消息认证码实现

// GenerateHmac 生成消息认证码
func GenerateHmac(plainText, key []byte) []byte {
	/*
	参数一 plainText：明文
	参数二 Key：密钥
	 */
	// 1. 创建哈希接口，需要指定使用的哈希算法和密钥
	myHash := hmac.New(sha1.New, key)
	// 2. 给哈希接口对象添加数据
	myHash.Write(plainText)
	// 3. 计算散列值
	hashText := myHash.Sum(nil)
	// 4. 如果是网络通信需要将计算后的 hash 值传输则需要执行下面编码步骤
	//hex.EncodeToString()
	return hashText
}

// VerifyHmac 校验消息认证码
func VerifyHmac(plainText, key,hashText []byte) bool {
	/*
	参数一 plainText：原始数据-明文
	参数二 key：密钥
	参数三 hashText: 生成的消息认证码
	 */
	// 1. 哈希接口，并执行使用的哈希算法阿和密钥
	myHash := hmac.New(sha1.New, key)
	// 2. 给哈希接口对象添加数据
	myHash.Write(plainText)
	// 3. 计算散列值
	hmac1 := myHash.Sum(nil)
	// 4. 与之前生成的消息验证码进行比较
	return hmac.Equal(hashText, hmac1)
}

// 测试
func main() {
	src := []byte("需要将要发送的数据进行哈希运算，将哈希值和原始数据一并发送")
	key := []byte("helloworld")
	hmac1 := GenerateHmac(src, key)
	bl := VerifyHmac(src, key, hmac1)
	fmt.Printf("校验结果：:%t\n",bl)
}
