package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	t.Log("开始Token生成测试...")

	// 测试场景1: 正常生成token
	t.Log("测试场景1: 正常生成token")
	userID := "123"
	token, err := GenerateToken(userID)
	if err != nil {
		t.Errorf("Token生成失败: %v", err)
	}
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
	t.Logf("成功生成token: %s", token)

	// 测试场景2: token唯一性
	t.Log("测试场景2: 验证token唯一性")
	token2, err := GenerateToken(userID)
	assert.Nil(t, err)
	assert.NotEqual(t, token, token2)
	t.Log("验证token唯一性成功")

	// 测试场景3: 空用户ID
	t.Log("测试场景3: 空用户ID处理")
	token3, err := GenerateToken("")
	assert.Nil(t, err)
	assert.NotEmpty(t, token3)
	t.Log("空用户ID处理成功")

	// 测试场景4: token长度验证
	t.Log("测试场景4: token长度验证")
	assert.GreaterOrEqual(t, len(token), 32, "token长度应该大于或等于32位")
	t.Log("token长度验证成功")

	t.Log("所有Token生成测试完成")
}

// 添加性能测试
func BenchmarkGenerateToken(b *testing.B) {
	b.Log("开始Token生成性能测试...")
	userID := "123"
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_, err := GenerateToken(userID)
		if err != nil {
			b.Errorf("Token生成失败: %v", err)
		}
	}
} 