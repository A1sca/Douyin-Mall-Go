package service

import (
	"context"
	"strconv"
	"testing"

	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/model"
	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"
	"github.com/stretchr/testify/assert"
)

func TestRegister_Run(t *testing.T) {
	t.Log("开始注册服务测试...")
	ctx := context.Background()
	s := NewRegisterService(ctx)

	// 测试场景1: 正常注册
	t.Log("测试场景1: 正常注册")
	req := &user.RegisterReq{
		Username: "testregister",
		Password: "testpassword",
		Email:    "testregister@example.com",
	}
	resp, err := s.Run(req)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.UserId)
	assert.NotEmpty(t, resp.Token)
	t.Log("正常注册测试通过")

	// 测试场景2: 重复用户名注册
	t.Log("测试场景2: 重复用户名注册")
	req = &user.RegisterReq{
		Username: "testregister",
		Password: "anotherpassword",
		Email:    "another@example.com",
	}
	resp, err = s.Run(req)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
	t.Logf("预期错误: %v", err)

	// 测试场景3: 空用户名注册
	t.Log("测试场景3: 空用户名注册")
	req = &user.RegisterReq{
		Username: "",
		Password: "testpassword",
		Email:    "test@example.com",
	}
	resp, err = s.Run(req)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
	t.Logf("预期错误: %v", err)

	// 测试场景4: 空密码注册
	t.Log("测试场景4: 空密码注册")
	req = &user.RegisterReq{
		Username: "testuser2",
		Password: "",
		Email:    "test@example.com",
	}
	resp, err = s.Run(req)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
	t.Logf("预期错误: %v", err)

	// 测试清理
	t.Log("清理测试数据...")
	user1, _ := model.GetByName(ctx, mysql.DB, "testregister")
	if user1 != nil {
		err = model.DeleteById(ctx, mysql.DB, strconv.FormatInt(int64(user1.ID), 10))
		assert.Nil(t, err)
	}
	t.Log("测试数据清理完成")

	t.Log("所有注册服务测试完成")
}
