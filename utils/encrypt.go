package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
	"log"
	"math/rand"
	"time"
)

type _encrypt struct{}

var Encryptor = new(_encrypt)

// 使用 scrypt 对密码进行加密生成一个哈希值
func (*_encrypt) ScryptHash(password string) string {
	const KeyLen = 10
	// 随便写
	salt := []byte{13, 14, 5, 2, 0, 55, 22, 11}

	hashPwd, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, KeyLen)
	if err != nil {
		log.Fatal("加密失败: ", err)
	}
	return base64.StdEncoding.EncodeToString(hashPwd)
}

// 使用 scrypt 对比 明文密码 和 数据库中哈希值
func (c *_encrypt) ScryptCheck(password, hash string) bool {
	return c.ScryptHash(password) == hash
}

// 使用 bcrypt 对密码进行加密生成一个哈希值
func (*_encrypt) BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// 使用 bcrypt 对比 明文密码 和 数据库中哈希值
func (*_encrypt) BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// MD5 加密
func (*_encrypt) MD5(str string, b ...byte) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(b))
}

// 验证码
func (*_encrypt) ValidateCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

// TODO
func UUID() string {
	return ""
}
