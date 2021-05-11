package main

import "fmt"

// 测试文件

func main() {
	fmt.Println("des 加解密开始执行, CBC 模式......")
	key := []byte("1234abcd") // 秘钥，8字节
	src := []byte("测试加解密代码内容，测试加解密代码内容") // 需要加密的明文
	cipherText := desEncrypt(src, key) // 调用实现的函数加密
	plainText := desDecrypt(cipherText, key) // 调用实现的函数解密
	fmt.Println("DES 解密成功结果：", string(plainText))
	// AES 加解密测试代码
	fmt.Println("aes 加解密开始执行 ctr 模式......")
	key1 := []byte("1234abdd12345678") // 秘钥 16 字节
	cipherText1 := aesEncrypt(src, key1)
	plainText1 := aesDecrypt(cipherText1, key1)
	fmt.Println("aes 解密成功结果：", string(plainText1))
	// 3des 加解密测试代码
	fmt.Println("3des 加解密开始执行 cbc 模式......")
	key2 := []byte("12345678abcdefgh87654321") // 秘钥24 字节
	cipherText2 := des3Encrypt(src, key2)
	plainText2 := des3Decrypt(cipherText2, key2)
	fmt.Println("3des 解密成功结果：", string(plainText2))
}