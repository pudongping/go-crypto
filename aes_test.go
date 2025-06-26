package go_crypto

import (
	"fmt"
	"testing"
)

func TestAESCBCEncrypt(t *testing.T) {
	plaintext := "hello world! My name is Alex Pu"
	// key 必须是 16(AES-128)、24(AES-192) 或 32(AES-256) 字节的 AES 密钥
	// key := "1234567890123456"
	// key := "123456789012345678901234"
	key := "12345678901234567890123456789012"
	fmt.Println("原文 ==> ", plaintext)
	ciphertext, err := AESCBCEncrypt(plaintext, key)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("密文 ==> ", ciphertext)
	plaintext2, err := AESCBCDecrypt(ciphertext, key)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("解密 ==> ", plaintext2)

	if plaintext != plaintext2 {
		t.Error("TestCBCEncrypt failed")
	}

}

func TestAESECBEncrypt(t *testing.T) {
	plaintext := "hello world! My name is Alex Pu"
	key := "1234567890123456"
	// key := "123456789012345678901234"
	// key := "12345678901234567890123456789012"
	fmt.Println("原文 ==> ", plaintext)
	ciphertext, err := AESECBEncrypt(plaintext, key)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("密文 ==> ", ciphertext)
	plaintext2, err := AESECBDecrypt(ciphertext, key)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("解密 ==> ", plaintext2)

	if plaintext != plaintext2 {
		t.Error("TestECBEncrypt failed")
	}

}

func TestAESCTREncrypt(t *testing.T) {
	plaintext := "hello world! My name is Alex Pu"
	// key := "1234567890123456"
	// key := "123456789012345678901234"
	key := "12345678901234567890123456789012"
	fmt.Println("原文 ==> ", plaintext)
	ciphertext, err := AESCTREncrypt(plaintext, key)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("密文 ==> ", ciphertext)
	plaintext2, err := AESCTRDecrypt(ciphertext, key)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("解密 ==> ", plaintext2)

	if plaintext != plaintext2 {
		t.Error("TestCTREncrypt failed")
	}

}

func TestAESCFBEncrypt(t *testing.T) {
	plaintext := "hello world! My name is Alex Pu"
	// key := "1234567890123456"
	// key := "123456789012345678901234"
	key := "12345678901234567890123456789012"
	fmt.Println("原文 ==> ", plaintext)
	ciphertext, err := AESCFBEncrypt(plaintext, key)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("密文 ==> ", ciphertext)
	plaintext2, err := AESCFBDecrypt(ciphertext, key)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("解密 ==> ", plaintext2)

	if plaintext != plaintext2 {
		t.Error("TestCFBEncrypt failed")
	}

}

func TestAESOFBEncrypt(t *testing.T) {
	plaintext := `{"code":200,"msg":"ok","data":{"id":1,"name":"张三","age":18}`
	// key := "1234567890123456"
	key := "123456789012345678901234"
	// key := "12345678901234567890123456789012"
	fmt.Println("原文 ==> ", plaintext)
	ciphertext, err := AESOFBEncrypt(plaintext, key)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("密文 ==> ", ciphertext)
	plaintext2, err := AESOFBDecrypt(ciphertext, key)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("解密 ==> ", plaintext2)

	if plaintext != plaintext2 {
		t.Error("TestOFBEncrypt failed")
	}
}

func TestAESGCMEncrypt(t *testing.T) {
	plaintext := "hello world! My name is Alex Pu"
	// key 支持三种密钥长度
	// key := "1234567890123456" // 16字节密钥 (AES-128-GCM)
	// key := "123456789012345678901234" // 24字节密钥 (AES-192-GCM)
	key := "12345678901234567890123456789012" // 32字节密钥 (AES-256-GCM)
	fmt.Println("原文 ==> ", plaintext)
	ciphertext, err := AESGCMEncrypt(plaintext, key)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("密文 ==> ", ciphertext)

	plaintext2, err := AESGCMDecrypt(ciphertext, key)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("解密 ==> ", plaintext2)

	if plaintext != plaintext2 {
		t.Error("TestAESGCMEncrypt failed")
	}
}
