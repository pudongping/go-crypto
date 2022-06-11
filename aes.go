package go_crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
)

func AESECBEncrypt(plaintext, key string) (string, error) {
	data := []byte(plaintext)
	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	length := (len(data) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, data)
	pad := byte(len(plain) - len(data))
	for i := len(data); i < len(plain); i++ {
		plain[i] = pad
	}
	cipher := make([]byte, len(plain))

	for bs, be := 0, block.BlockSize(); bs <= len(data); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Encrypt(cipher[bs:be], plain[bs:be])
	}

	return hex.EncodeToString(cipher), nil
}

func AESECBDecrypt(ciphertext, key string) (string, error) {
	encrypted, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	decrypted := make([]byte, len(encrypted))

	for bs, be := 0, block.BlockSize(); bs < len(encrypted); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return string(decrypted[:trim]), nil
}

func AESCBCEncrypt(plaintext, key string) (string, error) {
	data := []byte(plaintext)
	k := []byte(key)
	// 分组密钥
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	// 获取密钥块的长度
	blockSize := block.BlockSize()
	// 补充码
	data = PKCS7Padding(data, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	crypted := make([]byte, len(data))
	// 加密
	blockMode.CryptBlocks(crypted, data)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func AESCBCDecrypt(ciphertext, key string) (string, error) {
	cipherByte, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	k := []byte(key)
	// 分组密钥
	block, _ := aes.NewCipher(k)
	// 获取密钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	origData := make([]byte, len(cipherByte))
	// 解密
	blockMode.CryptBlocks(origData, cipherByte)
	// 去补码
	origData = PKCS7UnPadding(origData)
	return string(origData), nil
}

func AESCTREncrypt(plaintext, key string) (string, error) {

	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}

	// 创建分组模式
	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	stream := cipher.NewCTR(block, iv)

	// 加密
	dst := make([]byte, len(plaintext))
	stream.XORKeyStream(dst, []byte(plaintext))

	return string(dst), nil
}

func AESCTRDecrypt(ciphertext, key string) (string, error) {
	return AESCTREncrypt(ciphertext, key)
}

func AESCFBEncrypt(plaintext, key string) (string, error) {
	data := []byte(plaintext)
	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}

	encrypted := make([]byte, len(data)+aes.BlockSize)
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], data)

	return hex.EncodeToString(encrypted), nil
}

func AESCFBDecrypt(ciphertext, key string) (string, error) {
	encrypted, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}

	if len(encrypted) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)

	return string(encrypted), nil
}

func AESOFBEncrypt(plaintext, key string) (string, error) {
	data := PKCS7Padding([]byte(plaintext), aes.BlockSize)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	crypted := make([]byte, len(data)+aes.BlockSize)
	iv := crypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(crypted[aes.BlockSize:], data)
	return hex.EncodeToString(crypted), nil
}

func AESOFBDecrypt(ciphertext, key string) (string, error) {
	encrypted, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]
	if len(encrypted)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	origData := make([]byte, len(encrypted))
	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(origData, encrypted)

	return string(PKCS7UnPadding(origData)), nil
}

// PKCS7Padding 补码
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	// 判断缺少几位长度，最少为 1，最多为 blockSize
	padding := blockSize - len(ciphertext)%blockSize
	// 补足位数，把切片 []byte{byte(padding)} 复制 padding 个
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding 去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	if length == 0 {
		return nil
	}
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func generateKey(key []byte, length int) []byte {
	genKey := make([]byte, length)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
