package service

import (
	"context"
	"testing"
	"strconv"
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/model"
	"github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"
	"github.com/stretchr/testify/assert"
)

func TestDelete_Run(t *testing.T) {
	t.Log("=== 开始测试用户删除服务 ===")

	// 初始化上下文和服务
	t.Log("初始化测试环境...")
	ctx := context.Background()
	s := NewDeleteService(ctx)

	// 测试删除不存在的用户
	t.Log("\n测试场景1: 删除不存在的用户")
	req := &user.DeleteReq{
		UserId: "99999",
	}
	t.Logf("请求参数: %+v", req)
	resp, err := s.Run(req)
	t.Logf("响应: %+v, 错误: %v", resp, err)
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
	assert.Nil(t, resp)
	t.Log("预期的错误测试通过")

	// 创建测试用户
	t.Log("\n测试场景2: 创建测试用户")
	testUser := &model.User{
		Username: "deletetest",
		Email:    "deletetest@example.com",
	}
	t.Logf("创建用户信息: %+v", testUser)
	err = model.Create(ctx, mysql.DB, testUser)
	assert.NoError(t, err)
	t.Logf("测试用户创建成功，ID=%d", testUser.ID)

	// 测试删除存在的用户
	t.Log("\n测试场景3: 删除存在的用户")
	req = &user.DeleteReq{
		UserId: strconv.FormatUint(uint64(testUser.ID), 10),
	}
	t.Logf("请求参数: %+v", req)
	resp, err = s.Run(req)
	t.Logf("响应: %+v, 错误: %v", resp, err)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	t.Log("用户删除成功")

	// 验证用户确实被删除
	t.Log("\n测试场景4: 验证用户删除结果")
	deletedUser, err := model.GetById(ctx, mysql.DB, strconv.FormatUint(uint64(testUser.ID), 10))
	t.Logf("查询结果: %+v, 错误: %v", deletedUser, err)
	assert.NoError(t, err)
	assert.Nil(t, deletedUser)
	t.Log("验证通过：用户已被成功删除")

	t.Log("=== 用户删除服务测试完成 ===")
}
