package utils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmptyPassword = errors.New("password cannot be empty")
)

// HashPassword 对密码进行哈希处理
func HashPassword(password string) (string, error) {
	// 添加空密码检查
	if password == "" {
		return "", ErrEmptyPassword
	}
	
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPassword 验证密码
func CheckPassword(password, hash string) bool {
	if password == "" || hash == "" {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
} 