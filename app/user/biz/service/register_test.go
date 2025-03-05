package service

import (
	"context"
	"testing"

	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/model"
	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"
	"github.com/stretchr/testify/assert"
)

func TestRegister_Run(t *testing.T) {
	t.Log("=== 开始测试用户注册服务 ===")

	// 初始化数据库连接
	t.Log("初始化数据库连接...")
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)

	// 清理可能存在的测试数据
	_, err := model.GetByName(ctx, mysql.DB, "testuser1")
	if err == nil {
		err = model.DeleteByName(ctx, mysql.DB, "testuser1")
		assert.Nil(t, err)
	}

	// 测试场景1：正常注册
	t.Log("\n测试场景1: 正常注册流程")
	req1 := &user.RegisterReq{
		Username: "testuser1",
		Password: "password123",
		Email:    "testuser1@example.com",
	}
	t.Logf("请求参数: %+v", req1)

	resp1, err := s.Run(req1)
	assert.Nil(t, err)
	assert.NotNil(t, resp1)
	assert.NotEmpty(t, resp1.UserId)
	assert.NotEmpty(t, resp1.Token)

	// 测试场景2：重复注册
	t.Log("\n测试场景2: 重复注册测试")
	req2 := &user.RegisterReq{
		Username: "testuser1", // 使用相同的用户名
		Password: "password456",
		Email:    "testuser2@example.com",
	}
	t.Logf("请求参数: %+v", req2)

	_, err = s.Run(req2)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "username already exists")

	// 测试场景3：参数验证
	t.Log("\n测试场景3: 参数验证测试")
	req3 := &user.RegisterReq{
		Email: "invalid@example.com",
	}
	t.Logf("请求参数: %+v", req3)

	_, err = s.Run(req3)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "username or password cannot be empty")

	// 清理测试数据
	t.Log("清理测试数据...")
	err = model.DeleteByName(ctx, mysql.DB, "testuser1")
	assert.Nil(t, err)
	t.Log("测试数据清理完成")

	t.Log("=== 用户注册服务测试完成 ===")
}
