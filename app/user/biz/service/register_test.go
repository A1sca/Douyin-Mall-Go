package service

import (
	"context"
	"testing"
	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"
	"github.com/stretchr/testify/assert"
)

func TestRegister_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRegisterService(ctx)

	// 测试正常注册
	req := &user.RegisterReq{
		Username: "testuser",
		Password: "testpass123",
		Email:    "test@example.com",
	}
	resp, err := s.Run(req)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.UserId)
	assert.NotEmpty(t, resp.Token)

	// 测试重复用户名
	resp, err = s.Run(req)
	assert.NotNil(t, err)
	assert.Nil(t, resp)

	// 测试空用户名
	req.Username = ""
	resp, err = s.Run(req)
	assert.NotNil(t, err)
	assert.Nil(t, resp)

	// 测试空密码
	req.Username = "testuser2"
	req.Password = ""
	resp, err = s.Run(req)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}
