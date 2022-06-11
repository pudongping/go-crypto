package go_crypto

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestRSAEncrypt(t *testing.T) {
	// 可通过 openssl 生成公钥和私钥
	// 在当前目录下执行以下命令，即可生成 `rsa_private_key.pem` 和 `rsa_public_key.pem` 文件
	// openssl genrsa -out rsa_private_key.pem 1024
	// openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
	privateKey, err := ioutil.ReadFile("./rsa_private_key.pem")
	if err != nil {
		t.Error(err)
		return
	}

	publicKey, err := ioutil.ReadFile("./rsa_public_key.pem")
	if err != nil {
		t.Error(err)
		return
	}

	plaintext := "hello world"
	fmt.Println("原文 ==> ", plaintext)
	ciphertext, err := RSAEncrypt(publicKey, []byte(plaintext))
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("密文 base64 ==> ", base64.StdEncoding.EncodeToString(ciphertext))

	plaintext1, err := RSADecrypt(privateKey, ciphertext)
	fmt.Println("解密 ==> ", string(plaintext1))
	if err != nil {
		t.Error(err)
		return
	}

	if string(plaintext1) != plaintext {
		t.Error("RSAEncrypt and RSADecrypt error")
		return
	}

}
