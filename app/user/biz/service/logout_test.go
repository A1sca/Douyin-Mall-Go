package service

import (
	"context"
	"testing"
	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"
	"github.com/stretchr/testify/assert"
)

func TestLogout_Run(t *testing.T) {
	ctx := context.Background()
	s := NewLogoutService(ctx)

	// 测试正常登出
	req := &user.LogoutReq{
		UserId: "testuser123",
	}
	resp, err := s.Run(req)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.UserId, resp.UserId)

	// 测试空用户ID
	req = &user.LogoutReq{
		UserId: "",
	}
	resp, err = s.Run(req)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
} 