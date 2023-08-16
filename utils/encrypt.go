package utils

import "golang.org/x/crypto/bcrypt"

type _encrypt struct{}

var Encryptor = new(_encrypt)

// 使用bcrypt对密码进行加密生成一个哈希
func (*_encrypt) BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}
