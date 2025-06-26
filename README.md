# go-crypto
a encrypt and decrypt tools for golang


## 安装

```shell
go get -u github.com/pudongping/go-crypto
```

## 测试

```shell
go test ./ -v -count=1  
```

## 特性

### 1. 实现了 `AES` 加解密方法

> 推荐使用 `AES-GCM` 模式进行加解密。

- 电码本模式 （Electronic Codebook Book (ECB)）
- 密码分组链接模式（Cipher Block Chaining (CBC)）
- 计算器模式 （Counter (CTR)）
- 密码反馈模式 （Cipher FeedBack (CFB)）
- 输出反馈模式（Output FeedBack (OFB)）
- 伽罗瓦计数器模式（Galois/Counter Mode (GCM)）

### 2. 实现了 `RSA` 加解密方法

## AES

### GCM 模式

<details>
<summary>AES GCM 模式 Go、PHP 加解密示例 </summary>

go 加解密

```go
import "github.com/pudongping/go-crypto"

func main() {
	plaintext := "hello world! My name is Alex Pu"
	
    // key 支持三种密钥长度
    // key := "1234567890123456" // 16字节密钥 (AES-128-GCM)
    // key := "123456789012345678901234" // 24字节密钥 (AES-192-GCM)
    key := "12345678901234567890123456789012" // 32字节密钥 (AES-256-GCM)
	// 原文 ==>  hello world! My name is Alex Pu
    fmt.Println("原文 ==> ", plaintext)
    ciphertext, err := go_crypto.AESGCMEncrypt(plaintext, key)
    if err != nil {
        fmt.Println("出错啦！", err)
        return
    }
	// 密文 ==>  uvAupkvGwV/bHSIFL4rKtZe8gzHppM486pfcmwNSSBL0ZKkHwfjSD1QlpJe2bBPi/shdfzVSRf2Ke6s=
    fmt.Println("密文 ==> ", ciphertext)
    
    plaintext2, err := go_crypto.AESGCMDecrypt(ciphertext, key)
    if err != nil {
        fmt.Println("出错啦！", err)
        return
    }
	// 解密 ==>  hello world! My name is Alex Pu
    fmt.Println("解密 ==> ", plaintext2)	
}
```

php 加解密

```php
<?php

/**
 * AES-GCM 加密函数
 *
 * @param string $plaintext 明文字符串
 * @param string $key 密钥（16、24、32字节）
 * @param string $algo 加密算法名称，如 aes-128-gcm、aes-256-gcm
 * @return string           Base64 编码的 nonce + ciphertext + tag
 */
function aes_gcm_encrypt(string $plaintext, string $key, string $algo = 'aes-256-gcm'): string
{
    $ivLength = openssl_cipher_iv_length($algo);
    $iv = random_bytes($ivLength); // nonce

    $tag = '';
    $ciphertext = openssl_encrypt(
        $plaintext,
        $algo,
        $key,
        OPENSSL_RAW_DATA,
        $iv,
        $tag,
        '',      // AAD（可选）
        16       // Tag 长度
    );

    if ($ciphertext === false) {
        throw new RuntimeException('加密失败');
    }

    // 输出格式：Base64(nonce + ciphertext + tag)
    return base64_encode($iv . $ciphertext . $tag);
}

/**
 * AES-GCM 解密函数
 *
 * @param string $cipherBase64 Base64 编码的 nonce + ciphertext + tag
 * @param string $key 密钥（16、24、32字节）
 * @param string $algo 加密算法名称，如 aes-128-gcm、aes-256-gcm
 * @return string              明文
 */
function aes_gcm_decrypt(string $cipherBase64, string $key, string $algo = 'aes-256-gcm'): string
{
    $cipherData = base64_decode($cipherBase64);
    $ivLength = openssl_cipher_iv_length($algo);

    if ($ivLength === false || strlen($cipherData) < ($ivLength + 16)) {
        throw new InvalidArgumentException('密文格式错误');
    }

    $iv = substr($cipherData, 0, $ivLength);
    $tag = substr($cipherData, -16);
    $ciphertext = substr($cipherData, $ivLength, -16);

    $plaintext = openssl_decrypt(
        $ciphertext,
        $algo,
        $key,
        OPENSSL_RAW_DATA,
        $iv,
        $tag,
        '' // AAD（可选）
    );

    if ($plaintext === false) {
        throw new RuntimeException('解密失败');
    }

    return $plaintext;
}


$text = "hello world! My name is Alex Pu";

// $key = '1234567890123456';
// $algo = 'AES-128-GCM';
// $s = 'YEfwbkvJwUtwcSlqRhKf5xrVzzn9r3ZWj8JU8LKZUouzgcvCKvJMxgx36g5hKGQZUBcKWZxAoXipVHs=';

// $key = '123456789012345678901234';
// $algo = 'AES-192-GCM';
// $s = 'An0/iQf0BOR+g2qN31XAqcjV5Esp5HjkTO6Zy9pAMZ+rbIh0VPANqr1dP8S9t0bM51hByqqI3vkqLAc=';

$key = '12345678901234567890123456789012';
$algo = 'AES-256-GCM';
$s = 'SVpsIZ0LOGCN4XBaeALpxtXGecut+xPoAJHpcHBTXCRewkVvFCslxn1+T85tSAsD078O+SlpcxGQJ44=';

$a = aes_gcm_decrypt($s, $key, $algo);

// string(31) "hello world! My name is Alex Pu"
var_dump($a);

$aa = aes_gcm_encrypt($text, $key, $algo);
$bb = aes_gcm_decrypt($aa, $key, $algo);
var_dump($bb);
```


</details>

---

### CBC 模式


<details>
<summary>go 加密，php 解密（AES-128-CBC） </summary>

go 加密

```go
import "github.com/pudongping/go-crypto"

func main() {
    plaintext := "hello world! My name is Alex Pu"
	// 密钥字节长度必须为 16 个字节
    key := "1234567890123456"
	
    ciphertext, err := go_crypto.AESCBCEncrypt(plaintext, key)
    if err != nil {
        fmt.Println("出错啦！", err)
    }
	
    // output is: BRK08I0OYOoFwhgIBT1qjFywFkLADdeLQfVZM7CPKJ8=
    fmt.Println(ciphertext)
}

```

php 解密

```php

$key = '1234567890123456';
$iv = mb_substr($key, 0, 16);
$s = 'BRK08I0OYOoFwhgIBT1qjFywFkLADdeLQfVZM7CPKJ8=';

$str = base64_decode($s);
$decrypted = openssl_decrypt($str, 'AES-128-CBC', $key, OPENSSL_RAW_DATA, $iv);
if (!$decrypted) {
    echo '解密失败' . PHP_EOL;
} else {
    // output is: hello world! My name is Alex Pu
    echo($decrypted) . PHP_EOL;
}

```

</details>

---

<details>
<summary>php 加密，go 解密（AES-128-CBC） </summary>

php 加密

```php

$string = 'hello world! alex';
$key = '1234567890123456';
$iv = mb_substr($key, 0, 16);

$encrypted = openssl_encrypt($string, 'AES-128-CBC', $key, OPENSSL_RAW_DATA, $iv);
$s = base64_encode($encrypted);

// output is: YAZkDJYi7e9O09TRNvUf+6sFMlI8AQvZ5GVU+xJIuOc=
echo $s . PHP_EOL;

```

go 解密

```go
import "github.com/pudongping/go-crypto"

func main() {
    ciphertext := "YAZkDJYi7e9O09TRNvUf+6sFMlI8AQvZ5GVU+xJIuOc="
    key := "1234567890123456"
    
    plaintext, err := go_crypto.AESCBCDecrypt(ciphertext, key)
    if err != nil {
        fmt.Println("出错啦！", err)
    }
	
	// output is: 解密 ==>  hello world! alex
    fmt.Println("解密 ==> ", plaintext)
}

```

</details>

---

<details>
<summary>go 加密，php 解密（AES-192-CBC） </summary>

go 加密

```go
import "github.com/pudongping/go-crypto"

func main() {
    plaintext := "hello world! My name is Alex Pu"
	// 密钥字节长度必须为 24 个字节
    key := "123456789012345678901234"
	
    ciphertext, err := go_crypto.AESCBCEncrypt(plaintext, key)
    if err != nil {
        fmt.Println("出错啦！", err)
    }
	
    // output is: ebH1cP6XhScEg6c1PEElf/OnaJDOQ5hgUcSBg4+wrU4=
    fmt.Println(ciphertext)
}

```

php 解密

```php

$key = '123456789012345678901234';
$iv = mb_substr($key, 0, 16);
$s = 'ebH1cP6XhScEg6c1PEElf/OnaJDOQ5hgUcSBg4+wrU4=';

$str = base64_decode($s);
$decrypted = openssl_decrypt($str, 'AES-192-CBC', $key, OPENSSL_RAW_DATA, $iv);
if (!$decrypted) {
    echo '解密失败' . PHP_EOL;
} else {
    // output is: hello world! My name is Alex Pu
    echo($decrypted) . PHP_EOL;
}

```

</details>

---

<details>
<summary>php 加密，go 解密（AES-192-CBC） </summary>

php 加密

```php

$string = 'hello world! alex';
$key = '123456789012345678901234';
$iv = mb_substr($key, 0, 16);

$encrypted = openssl_encrypt($string, 'AES-192-CBC', $key, OPENSSL_RAW_DATA, $iv);
$s = base64_encode($encrypted);

// output is: bBqpDd0iaOh3eKKgvsAWw+OnH5v8fBvb/8n+hE9YfbY=
echo $s . PHP_EOL;

```

go 解密

```go
import "github.com/pudongping/go-crypto"

func main() {
    ciphertext := "bBqpDd0iaOh3eKKgvsAWw+OnH5v8fBvb/8n+hE9YfbY="
    key := "123456789012345678901234"
    
    plaintext, err := go_crypto.AESCBCDecrypt(ciphertext, key)
    if err != nil {
        fmt.Println("出错啦！", err)
    }
	
	// output is: 解密 ==>  hello world! alex
    fmt.Println("解密 ==> ", plaintext)
}

```

</details>

---

<details>
<summary>go 加密，php 解密（AES-256-CBC） </summary>

go 加密

```go
import "github.com/pudongping/go-crypto"

func main() {
    plaintext := "hello world! My name is Alex Pu"
	// 密钥字节长度必须为 32 个字节
    key := "12345678901234567890123456789012"
	
    ciphertext, err := go_crypto.AESCBCEncrypt(plaintext, key)
    if err != nil {
        fmt.Println("出错啦！", err)
    }
	
    // output is: wwtsrRbifJPxKyxhpWk4WZ2RbzhEwES04tZBjAaC4pA=
    fmt.Println(ciphertext)
}

```

php 解密

```php

$key = '12345678901234567890123456789012';
$iv = mb_substr($key, 0, 16);
$s = 'wwtsrRbifJPxKyxhpWk4WZ2RbzhEwES04tZBjAaC4pA=';

$str = base64_decode($s);
$decrypted = openssl_decrypt($str, 'AES-256-CBC', $key, OPENSSL_RAW_DATA, $iv);
if (!$decrypted) {
    echo '解密失败' . PHP_EOL;
} else {
    // output is: hello world! My name is Alex Pu
    echo($decrypted) . PHP_EOL;
}

```

</details>

---

<details>
<summary>php 加密，go 解密（AES-256-CBC） </summary>

php 加密

```php

$string = 'hello world! alex';
$key = '12345678901234567890123456789012';
$iv = mb_substr($key, 0, 16);

$encrypted = openssl_encrypt($string, 'AES-256-CBC', $key, OPENSSL_RAW_DATA, $iv);
$s = base64_encode($encrypted);

// output is: RXqMpMoznbkuQFgU4ZXWEkmF14bIHbw2ncvmYqvN/sI=
echo $s . PHP_EOL;

```

go 解密

```go
import "github.com/pudongping/go-crypto"

func main() {
    ciphertext := "RXqMpMoznbkuQFgU4ZXWEkmF14bIHbw2ncvmYqvN/sI="
    key := "12345678901234567890123456789012"
    
    plaintext, err := go_crypto.AESCBCDecrypt(ciphertext, key)
    if err != nil {
        fmt.Println("出错啦！", err)
    }
	
	// output is: 解密 ==>  hello world! alex
    fmt.Println("解密 ==> ", plaintext)
}

```

</details>

---


### ECB 模式

> `AES-192-ECB` 和 `AES-256-ECB` 加密方式与 `AES-128-ECB` 加密方式大致一样，请结合以上提供的 `AES-CBC` 相关代码来编写。

<details>
<summary>go 加密，php 解密（AES-128-ECB） </summary>

go 加密

```go
import "github.com/pudongping/go-crypto"

func main() {
    plaintext := "hello world! My name is Alex Pu"
	// 密钥字节长度必须为 16 个字节
    key := "1234567890123456"
	
    ciphertext, err := go_crypto.AESECBEncrypt(plaintext, key)
    if err != nil {
        fmt.Println("出错啦！", err)
    }
	
    // output is: sRFeHhndretZFZE9/7WdGuGw1QYl8l/IlI1XEtpVzxI=
    fmt.Println(ciphertext)
}

```

php 解密

```php

$key = '1234567890123456';
$s = 'sRFeHhndretZFZE9/7WdGuGw1QYl8l/IlI1XEtpVzxI=';

$str = base64_decode($s);
$decrypted = openssl_decrypt($str, 'AES-128-ECB', $key, OPENSSL_RAW_DATA);
if (!$decrypted) {
    echo '解密失败' . PHP_EOL;
} else {
    // output is: hello world! My name is Alex Pu
    echo($decrypted) . PHP_EOL;
}

```

</details>

---

<details>
<summary>php 加密，go 解密（AES-128-ECB） </summary>

php 加密

```php

$string = 'hello world! alex';
$key = '1234567890123456';

$encrypted = openssl_encrypt($string, 'AES-128-ECB', $key, OPENSSL_RAW_DATA);
$s = base64_encode($encrypted);

// output is: 7LVm2y6R7E+Hj3mIRlHDbOWZPsz+Vvb2zOkt6htAttc=
echo $s . PHP_EOL;

```

go 解密

```go
import "github.com/pudongping/go-crypto"

func main() {
    ciphertext := "7LVm2y6R7E+Hj3mIRlHDbOWZPsz+Vvb2zOkt6htAttc="
    key := "1234567890123456"
    
    plaintext, err := go_crypto.AESECBDecrypt(ciphertext, key)
    if err != nil {
        fmt.Println("出错啦！", err)
    }
	
	// output is: 解密 ==>  hello world! alex
    fmt.Println("解密 ==> ", plaintext)
}

```

</details>

---

## RSA

<details>
<summary>rsa 加解密 </summary>

```go

package main

import (
	"fmt"

	"github.com/pudongping/go-crypto"
)

func main() {
	var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQC8haHPNLshJPplmf5jOh6fVgtLnRNOJh4qhOZY0YgwuIRQ+lOv
5f+kypKaU9YuGWQk5zb/6wVtxKZ3lApzqqpQqndtx/7cMWOzPIyIoJYenqUuaZ2m
JR0OLkSMPnncGMjaVfgKB07cl6q6l2xsR6e/WIwu+wxya6bBWqpd2K4/iQIDAQAB
AoGAOJaEM4ZIhXJsFwNacL6JYpqZWWEU4mhetbSe0BpVHwrD7Oq3JB2N1xFXYCu6
JqxIeytMPxV007Yrock2OfIInRmF8UK6H0bh/o+miAXHU2ILTWdOpdBM7KXIdj8k
8uxy8Dyp7PFfxDJZugx0Xq+VBhx+dPQV9eVAUoVX417ZVAECQQDyh8nKL4Knbaal
Js/hEgxiN24G78Y/9z/eiwGAzB8pCq4jPFjRxpWVXwUTUvphuXyNDKv7fKhhb0KP
uAeh6F0RAkEAxv34nzyIRSEu0kF/HuMAhRzeBaiaQs8FJAmMTsNC54vT65mGK3oL
2TNKacfjUuIBmpy/PkztL7f3e8su4Lga+QJAQFMvfAaZ2ppEhq0CmalVy370Gbbi
l/iV4gpwPZ7l7Na+VO4eiJTo+5MWH4f0jJLLrAVeX/cyGZWOy9t9er3MYQJBAK3N
e47RVxBSKEv6auZ2TKj3CrCUj4/Us1/Deyn7//6vMXyxDfABXEHsO41iyhlFTzNU
vvDSTcQFAYK9d4dNJjkCQQCxM1bTNieIiHuywQtNVD4EYGbu8T+holpSLpUiRQFA
cYgvv3oqJmElg6TNoM3n3K7rsmiAO24exPEkXQz0oMRM
-----END RSA PRIVATE KEY-----
`)

	var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC8haHPNLshJPplmf5jOh6fVgtL
nRNOJh4qhOZY0YgwuIRQ+lOv5f+kypKaU9YuGWQk5zb/6wVtxKZ3lApzqqpQqndt
x/7cMWOzPIyIoJYenqUuaZ2mJR0OLkSMPnncGMjaVfgKB07cl6q6l2xsR6e/WIwu
+wxya6bBWqpd2K4/iQIDAQAB
-----END PUBLIC KEY-----
`)

	plaintext := "hello world"
	fmt.Println("原文 ==> ", plaintext)
	ciphertext, err := go_crypto.RSAEncrypt(publicKey, []byte(plaintext))
	if err != nil {
		fmt.Println(err)
		return
	}

	plaintext1, err := go_crypto.RSADecrypt(privateKey, ciphertext)
	fmt.Println("解密 ==> ", string(plaintext1))
	if err != nil {
		fmt.Println(err)
		return
	}

}

```

</details>

---