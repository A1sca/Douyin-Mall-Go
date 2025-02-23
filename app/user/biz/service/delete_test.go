package service

import (
	"context"
	"testing"

	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/model"
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/dal/mysql"
	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"
	"github.com/stretchr/testify/assert"
)

func TestDelete_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteService(ctx)

	// 先创建一个测试用户
	testUser := &model.User{
		Username: "deletetest",
		Email:    "deletetest@example.com",
		Phone:    "13800138001",
	}
	err := model.Create(ctx, mysql.DB, testUser)
	assert.Nil(t, err)

	// 测试正常删除
	req := &user.DeleteReq{UserId: testUser.ID}
	resp, err := s.Run(req)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	// 测试删除不存在的用户
	req = &user.DeleteReq{UserId: 99999}
	resp, err = s.Run(req)
	assert.NotNil(t, err)
}

func TestDeleteService_DeleteUser(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteService(ctx)

	// 先创建一个测试用户
	testUser := &model.User{
		Username: "deletetest2",
		Email:    "deletetest2@example.com",
		Phone:    "13800138002",
	}
	err := model.Create(ctx, mysql.DB, testUser)
	assert.Nil(t, err)

	// 测试正常删除
	err = s.DeleteUser(ctx, int64(testUser.ID))
	assert.Nil(t, err)

	// 测试无效的用户ID
	err = s.DeleteUser(ctx, 0)
	assert.NotNil(t, err)

	// 测试删除不存在的用户
	err = s.DeleteUser(ctx, 99999)
	assert.NotNil(t, err)
}
