package utils

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	t.Log("开始测试token生成功能...")

	// 测试正常生成token
	t.Log("测试场景1: 正常生成token")
	token, err := GenerateToken("123", "testuser")
	if err != nil {
		t.Errorf("生成token失败: %v", err)
	}
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
	t.Logf("成功生成token: %s", token)

	// 测试解析生成的token
	t.Log("测试场景2: 解析刚生成的token")
	claims, err := ParseToken(token)
	if err != nil {
		t.Errorf("解析token失败: %v", err)
	}
	assert.Nil(t, err)
	assert.Equal(t, "123", claims.UserID)
	assert.Equal(t, "testuser", claims.Username)
	t.Logf("成功解析token, 用户ID: %s, 用户名: %s", claims.UserID, claims.Username)
	t.Logf("token详细信息: 过期时间: %v, 签发时间: %v, 生效时间: %v", 
		claims.ExpiresAt.Time, 
		claims.IssuedAt.Time, 
		claims.NotBefore.Time)
}

func TestParseToken(t *testing.T) {
	t.Log("开始测试token解析功能...")

	// 测试正常token
	t.Log("测试场景1: 解析有效token")
	token, _ := GenerateToken("123", "testuser")
	t.Logf("生成测试token: %s", token)
	claims, err := ParseToken(token)
	assert.Nil(t, err)
	assert.Equal(t, "123", claims.UserID)
	assert.Equal(t, "testuser", claims.Username)
	t.Log("成功解析有效token")

	// 测试过期token
	t.Log("测试场景2: 解析过期token")
	expiredClaims := JWTClaims{
		UserID:   "123",
		Username: "testuser",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Hour)), // 1小时前过期
		},
	}
	expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, expiredClaims)
	expiredTokenString, _ := expiredToken.SignedString(jwtSecret)
	t.Logf("生成过期token: %s", expiredTokenString)
	_, err = ParseToken(expiredTokenString)
	assert.Equal(t, TokenExpired, err)
	t.Logf("预期错误: %v", err)

	// 测试未生效token
	t.Log("测试场景3: 解析未生效token")
	notValidClaims := JWTClaims{
		UserID:   "123",
		Username: "testuser",
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Hour)), // 1小时后生效
		},
	}
	notValidToken := jwt.NewWithClaims(jwt.SigningMethodHS256, notValidClaims)
	notValidTokenString, _ := notValidToken.SignedString(jwtSecret)
	t.Logf("生成未生效token: %s", notValidTokenString)
	_, err = ParseToken(notValidTokenString)
	assert.Equal(t, TokenNotValidYet, err)
	t.Logf("预期错误: %v", err)

	// 测试格式错误的token
	t.Log("测试场景4: 解析格式错误的token")
	invalidToken := "invalid.token.string"
	t.Logf("使用无效token: %s", invalidToken)
	_, err = ParseToken(invalidToken)
	assert.Equal(t, TokenMalformed, err)
	t.Logf("预期错误: %v", err)

	t.Log("所有token解析测试完成")
}

// 添加性能测试
func BenchmarkGenerateToken(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GenerateToken("123", "testuser")
	}
}

func BenchmarkParseToken(b *testing.B) {
	token, _ := GenerateToken("123", "testuser")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParseToken(token)
	}
}