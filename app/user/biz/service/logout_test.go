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
	t.Logf("[测试场景1] 正常用户登出: %+v", req)
	resp, err := s.Run(req)
	t.Logf("[测试结果1] 响应: %+v, 错误: %v", resp, err)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.UserId, resp.UserId)

	// 测试空用户ID
	req = &user.LogoutReq{
		UserId: "",
	}
	t.Logf("[测试场景2] 空用户ID登出: %+v", req)
	resp, err = s.Run(req)
	t.Logf("[测试结果2] 响应: %+v, 错误: %v", resp, err)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}