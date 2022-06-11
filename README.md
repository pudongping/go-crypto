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

- 电码本模式 （Electronic Codebook Book (ECB)）
- 密码分组链接模式（Cipher Block Chaining (CBC)）
- 计算器模式 （Counter (CTR)）
- 密码反馈模式 （Cipher FeedBack (CFB)）
- 输出反馈模式（Output FeedBack (OFB)）

### 2. 实现了 `RSA` 加解密方法

## AES

### CBC 模式

- go 加密，php 解密（AES-128-CBC）

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

- php 加密，go 解密（AES-128-CBC）

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

- go 加密，php 解密（AES-192-CBC）

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

- php 加密，go 解密（AES-192-CBC）

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

- go 加密，php 解密（AES-256-CBC）

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

- php 加密，go 解密（AES-256-CBC）

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

---

### ECB 模式

- go 加密，php 解密（AES-128-ECB）

> `AES-192-ECB` 和 `AES-256-ECB` 加密方式与 `AES-128-ECB` 加密方式大致一样，请结合以上提供的 `AES-CBC` 相关代码来编写。

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

- php 加密，go 解密（AES-128-ECB）

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

## RSA

- rsa 加解密

```go

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
	ciphertext, err := RSAEncrypt(publicKey, []byte(plaintext))
	if err != nil {
        fmt.Println(err)
        return
	}

	plaintext1, err := RSADecrypt(privateKey, ciphertext)
	fmt.Println("解密 ==> ", string(plaintext1))
	if err != nil {
        fmt.Println(err)
        return
	}

```