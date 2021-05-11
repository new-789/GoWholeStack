package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
)

// DES CBC 加密

/* 思路
编写填充函数，如果最后一个分组字节数不够，填充
字节数刚好合适，添加一个新的分组
填充的字节的值 == 缺少的字节的数
*/
func paddingLastGroup(plainText []byte, blockSize int) []byte {
	// 1. 通过取余求出最后一个组中剩余的字节数
	padNum := blockSize - len(plainText)%blockSize
	// 2. 创建新的切片，长度等于 padNum，每个字节值 byte(padNum)
	char := []byte{byte(padNum)} // 长度1
	// 切片创建，并初始化
	newPlain := bytes.Repeat(char, padNum)
	// 3. newPlain 切片追加到原始明文后边
	newText := append(plainText, newPlain...)
	return newText
}

// 去掉分组尾部填充的数据
func unPaddingLast(plainText []byte) []byte {
	// 1. 拿出切片中的最后一个字节
	length := len(plainText)
	lastChar := plainText[length-1] // byte 类型
	number := int(lastChar) // 尾部填充的字节个数
	return plainText[:length-number]
}

// des 加密
func desEncrypt(plainText, key []byte) []byte {
	// 1. 建立一个底层使用 des 的密码接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 明文填充
	newText := paddingLastGroup(plainText, block.BlockSize())
	// 3. 创建一个使用 CBC 分组的接口
	iv := []byte("12345678") // 指定的初始化向量
	blockMode := cipher.NewCBCEncrypter(block, iv)
	// 4. 加密, 该方法会将第二个参数的内容加密，并存储到第一个参数中(切片)
	cipherText := make([]byte, len(newText))
	blockMode.CryptBlocks(cipherText, newText) // 也可以使用下面语法
	//blockMode.CryptBlocks(newText, newText)
	return cipherText
}

// des 解密
func desDecrypt(cipherText, key []byte) []byte {
	// 1. 创建一个底层使用 des 的密码接口
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 创建一个使用 des 模式的解密接口
	iv := []byte("12345678") // 指定的初始化向量
	blockMode := cipher.NewCBCDecrypter(block, iv)
	// 3. 解密
	blockMode.CryptBlocks(cipherText, cipherText)
	// 4. cipherText 现在存储的是明文，需要删除加密时填充的尾部数据
	plainText := unPaddingLast(cipherText)
	return plainText
}

// =======================================================================================
// aes 加密，分组模式 ctr
func aesEncrypt(plainText, key []byte) []byte {
	// 1. 建立一个底层使用 aes 的密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 创建一个使用 CTR 分组的接口
	iv := []byte("12345678abcdefgh") //  iv 为初始化的随机数种子 16 字节长度
	stream := cipher.NewCTR(block, iv)
	// 3. 加密, 该方法会将第二个参数的内容加密，并存储到第一个参数中(切片)
	cipherText := make([]byte, len(plainText))
	// 按位异或操作加密
	stream.XORKeyStream(cipherText, plainText)
	return cipherText
}

// aes 解密
func aesDecrypt(cipherText, key []byte) []byte {
	// 1. 创建一个底层使用 aes 的密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 创建一个使用 ctr 模式的解密接口
	iv := []byte("12345678abcdefgh") //  iv 为初始化的随机数种子 16 字节长度
	stream := cipher.NewCTR(block, iv)
	// 3. 按位异或操作解密，该方法会将第二个参数的内容加密，并存储到第一个参数中(切片)
	stream.XORKeyStream(cipherText, cipherText)
	return cipherText
}

// ========================================================================================
// 3des 加密
func des3Encrypt(plainText, key []byte) []byte {
	// 1. 建立一个底层使用 des 的密码接口
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 明文填充
	newText := paddingLastGroup(plainText, block.BlockSize())
	// 3. 创建一个使用 CBC 分组的接口
	iv := []byte("12345678") // 指定的初始化向量
	blockMode := cipher.NewCBCEncrypter(block, iv)
	// 4. 加密, 该方法会将第二个参数的内容加密，并存储到第一个参数中(切片)
	cipherText := make([]byte, len(newText))
	blockMode.CryptBlocks(cipherText, newText) // 也可以使用下面语法
	//blockMode.CryptBlocks(newText, newText)
	return cipherText
}

// 3des 解密
func des3Decrypt(cipherText, key []byte) []byte {
	// 1. 创建一个底层使用 des 的密码接口
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}
	// 2. 创建一个使用 des 模式的解密接口
	iv := []byte("12345678") // 指定的初始化向量
	blockMode := cipher.NewCBCDecrypter(block, iv)
	// 3. 解密
	blockMode.CryptBlocks(cipherText, cipherText)
	// 4. cipherText 现在存储的是明文，需要删除加密时填充的尾部数据
	plainText := unPaddingLast(cipherText)
	return plainText
}