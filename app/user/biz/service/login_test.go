package service

import (
	"context"
	"testing"
	
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/model"
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/utils"
	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"
	"github.com/stretchr/testify/assert"
)

func TestLogin_Run(t *testing.T) {
	t.Log("开始登录服务测试...")
	ctx := context.Background()
	s := NewLoginService(ctx)

	// 测试准备：创建测试用户
	t.Log("准备测试数据: 创建测试用户")
	password := "testpassword"
	hashedPassword, _ := utils.HashPassword(password)
	testUser := &model.User{
		Username: "testuser",
		Password: hashedPassword,
		Email:    "test@example.com",
	}
	err := model.Create(ctx, mysql.DB, testUser)
	assert.Nil(t, err)
	t.Log("测试用户创建成功")

	// 测试场景1: 正常登录
	t.Log("测试场景1: 正常登录")
	req := &user.LoginReq{
		Username: "testuser",
		Password: "testpassword",
	}
	resp, err := s.Run(req)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	t.Log("正常登录测试通过")

	// 测试场景2: 空用户名登录
	t.Log("测试场景2: 空用户名登录")
	req = &user.LoginReq{
		Username: "",
		Password: "testpassword",
	}
	resp, err = s.Run(req)
	assert.NotNil(t, err)
	t.Logf("预期错误: %v", err)

	// 测试场景3: 空密码登录
	t.Log("测试场景3: 空密码登录")
	req = &user.LoginReq{
		Username: "testuser",
		Password: "",
	}
	resp, err = s.Run(req)
	assert.NotNil(t, err)
	t.Logf("预期错误: %v", err)

	// 测试场景4: 错误密码登录
	t.Log("测试场景4: 错误密码登录")
	req = &user.LoginReq{
		Username: "testuser",
		Password: "wrongpassword",
	}
	resp, err = s.Run(req)
	assert.NotNil(t, err)
	t.Logf("预期错误: %v", err)

	// 测试清理
	t.Log("清理测试数据...")
	err = model.DeleteById(ctx, mysql.DB, testUser.ID)
	assert.Nil(t, err)
	t.Log("测试数据清理完成")

	t.Log("所有登录服务测试完成")
}
