package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
)

var (
	ErrGenerateToken = errors.New("failed to generate token")
)

// GenerateToken 生成token
func GenerateToken(userID string) (string, error) {
	// 生成随机字节
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", ErrGenerateToken
	}
	
	// 使用base64编码
	token := base64.URLEncoding.EncodeToString(b)
	return token, nil
} 