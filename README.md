# go-crypto
a encrypt and decrypt tools for golang


## 安装

```go
go get github.com/pudongping/go-crypto
```

## 特性

### 实现了 `AES` 加密算法

1. 电码本模式 （Electronic Codebook Book (ECB)）
2. 密码分组链接模式（Cipher Block Chaining (CBC)）
3. 计算器模式 （Counter (CTR)）
4. 密码反馈模式 （Cipher FeedBack (CFB)）
5. 输出反馈模式（Output FeedBack (OFB)）

## AES

### CBC 模式

- go 加密，php 解密（AES-128-CBC）

加密

```go
import "github.com/pudongping/go-crypto"

func main() {
    plaintext := "hello world! My name is Alex Pu"
	// 密钥字节长度必须为 16 个字节
    key := "1234567890123456"
	
    ciphertext, err := go-crypto.AESCBCEncrypt(plaintext, key)
    if err != nil {
        fmt.Println("出错啦！", err)
    }
	
    // output is: BRK08I0OYOoFwhgIBT1qjFywFkLADdeLQfVZM7CPKJ8=
    fmt.Println(ciphertext)
}

```

解密

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

加密

```php

$string = 'hello world! alex';
$key = '1234567890123456';
$iv = mb_substr($key, 0, 16);

$encrypted = openssl_encrypt($string, 'AES-128-CBC', $key, OPENSSL_RAW_DATA, $iv);
$s = base64_encode($encrypted);

// output is: YAZkDJYi7e9O09TRNvUf+6sFMlI8AQvZ5GVU+xJIuOc=
echo $s . PHP_EOL;

```

解密

```go
import "github.com/pudongping/go-crypto"

func main() {
    ciphertext := "YAZkDJYi7e9O09TRNvUf+6sFMlI8AQvZ5GVU+xJIuOc="
    key := "1234567890123456"
    
    plaintext, err := go-crypto.AESCBCDecrypt(ciphertext, key)
    if err != nil {
        fmt.Println("出错啦！", err)
    }
	
	// output is: 解密 ==>  hello world! alex
    fmt.Println("解密 ==> ", plaintext)
}

```

- go 加密，php 解密（AES-192-CBC）

加密

```go
import "github.com/pudongping/go-crypto"

func main() {
    plaintext := "hello world! My name is Alex Pu"
	// 密钥字节长度必须为 24 个字节
    key := "123456789012345678901234"
	
    ciphertext, err := go-crypto.AESCBCEncrypt(plaintext, key)
    if err != nil {
        fmt.Println("出错啦！", err)
    }
	
    // output is: ebH1cP6XhScEg6c1PEElf/OnaJDOQ5hgUcSBg4+wrU4=
    fmt.Println(ciphertext)
}

```

解密

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

加密

```php

$string = 'hello world! alex';
$key = '123456789012345678901234';
$iv = mb_substr($key, 0, 16);

$encrypted = openssl_encrypt($string, 'AES-192-CBC', $key, OPENSSL_RAW_DATA, $iv);
$s = base64_encode($encrypted);

// output is: bBqpDd0iaOh3eKKgvsAWw+OnH5v8fBvb/8n+hE9YfbY=
echo $s . PHP_EOL;

```

解密

```go
import "github.com/pudongping/go-crypto"

func main() {
    ciphertext := "bBqpDd0iaOh3eKKgvsAWw+OnH5v8fBvb/8n+hE9YfbY="
    key := "123456789012345678901234"
    
    plaintext, err := go-crypto.AESCBCDecrypt(ciphertext, key)
    if err != nil {
        fmt.Println("出错啦！", err)
    }
	
	// output is: 解密 ==>  hello world! alex
    fmt.Println("解密 ==> ", plaintext)
}

```

- go 加密，php 解密（AES-256-CBC）

加密

```go
import "github.com/pudongping/go-crypto"

func main() {
    plaintext := "hello world! My name is Alex Pu"
	// 密钥字节长度必须为 32 个字节
    key := "12345678901234567890123456789012"
	
    ciphertext, err := go-crypto.AESCBCEncrypt(plaintext, key)
    if err != nil {
        fmt.Println("出错啦！", err)
    }
	
    // output is: wwtsrRbifJPxKyxhpWk4WZ2RbzhEwES04tZBjAaC4pA=
    fmt.Println(ciphertext)
}

```

解密

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

加密

```php

$string = 'hello world! alex';
$key = '12345678901234567890123456789012';
$iv = mb_substr($key, 0, 16);

$encrypted = openssl_encrypt($string, 'AES-256-CBC', $key, OPENSSL_RAW_DATA, $iv);
$s = base64_encode($encrypted);

// output is: RXqMpMoznbkuQFgU4ZXWEkmF14bIHbw2ncvmYqvN/sI=
echo $s . PHP_EOL;

```

解密

```go
import "github.com/pudongping/go-crypto"

func main() {
    ciphertext := "RXqMpMoznbkuQFgU4ZXWEkmF14bIHbw2ncvmYqvN/sI="
    key := "12345678901234567890123456789012"
    
    plaintext, err := go-crypto.AESCBCDecrypt(ciphertext, key)
    if err != nil {
        fmt.Println("出错啦！", err)
    }
	
	// output is: 解密 ==>  hello world! alex
    fmt.Println("解密 ==> ", plaintext)
}

```