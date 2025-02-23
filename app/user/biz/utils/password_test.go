package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPasswordHash(t *testing.T) {
	t.Log("开始密码哈希测试...")

	// 测试场景1: 正常密码哈希
	t.Log("测试场景1: 正常密码哈希")
	password := "testpassword123"
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("密码哈希失败: %v", err)
	}
	assert.Nil(t, err)
	assert.NotEmpty(t, hash)
	t.Logf("成功生成密码哈希: %s", hash)

	// 测试场景2: 正确密码验证
	t.Log("测试场景2: 正确密码验证")
	isValid := CheckPassword(password, hash)
	assert.True(t, isValid)
	t.Log("密码验证成功")

	// 测试场景3: 错误密码验证
	t.Log("测试场景3: 错误密码验证")
	isValid = CheckPassword("wrongpassword", hash)
	assert.False(t, isValid)
	t.Log("错误密码验证通过测试")

	// 测试场景4: 空密码处理
	t.Log("测试场景4: 空密码处理")
	hash, err = HashPassword("")
	assert.Equal(t, ErrEmptyPassword, err, "应该返回空密码错误")
	assert.Empty(t, hash)
	t.Log("空密码处理成功")

	// 测试场景5: 空密码验证
	t.Log("测试场景5: 空密码验证")
	isValid = CheckPassword("", hash)
	assert.False(t, isValid)
	t.Log("空密码验证测试通过")

	t.Log("所有密码哈希测试完成")
}

// 添加性能测试
func BenchmarkHashPassword(b *testing.B) {
	b.Log("开始密码哈希性能测试...")
	password := "testpassword123"
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_, err := HashPassword(password)
		if err != nil {
			b.Errorf("密码哈希失败: %v", err)
		}
	}
}

func BenchmarkCheckPassword(b *testing.B) {
	b.Log("开始密码验证性能测试...")
	password := "testpassword123"
	hash, _ := HashPassword(password)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		CheckPassword(password, hash)
	}
} 