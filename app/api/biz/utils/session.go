package utils

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"github.com/A1sca/Douyin-Mall-Go/app/frontend/biz/dal/redis"
)

// GenerateSessionID 生成随机的会话 ID
func GenerateSessionID() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// GetUserIDFromSession 从 session 获取用户 ID
func GetUserIDFromSession(ctx context.Context, sessionID string) (string, error) {
	userID, err := redis.RedisClient.Get(ctx, sessionID).Result()
	if err != nil {
		return "", err
	}
	return userID, nil
}
