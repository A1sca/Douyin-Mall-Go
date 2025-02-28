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

func TestGet_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetService(ctx)

	// 创建测试用户
	testUser := &model.User{
		Username: "gettest",
		Email:    "gettest@example.com",
		Phone:    "13800138003",
	}
	err := model.Create(ctx, mysql.DB, testUser)
	assert.Nil(t, err)

	// 测试正常获取
	req := &user.GetReq{UserId: strconv.FormatUint(uint64(testUser.ID), 10)}
	resp, err := s.Run(req)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	// 测试获取不存在的用户
	req = &user.GetReq{UserId: "99999"}
	resp, err = s.Run(req)
	assert.NotNil(t, err)
}

func TestGetService_GetUser(t *testing.T) {
	ctx := context.Background()
	s := NewGetService(ctx)

	// 创建测试用户
	testUser := &model.User{
		Username: "gettest2",
		Email:    "gettest2@example.com",
		Phone:    "13800138004",
	}
	err := model.Create(ctx, mysql.DB, testUser)
	assert.Nil(t, err)

	// 测试正常获取
	user, err := s.GetUser(ctx, int64(testUser.ID))
	assert.Nil(t, err)
	assert.Equal(t, testUser.Username, user.Username)

	// 测试无效的用户ID
	user, err = s.GetUser(ctx, 0)
	assert.NotNil(t, err)
	assert.Nil(t, user)

	// 测试不存在的用户
	user, err = s.GetUser(ctx, 99999)
	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestGetService_GetUserByUsername(t *testing.T) {
	ctx := context.Background()
	s := NewGetService(ctx)

	// 创建测试用户
	testUser := &model.User{
		Username: "gettest3",
		Email:    "gettest3@example.com",
		Phone:    "13800138005",
	}
	err := model.Create(ctx, mysql.DB, testUser)
	assert.Nil(t, err)

	// 测试正常获取
	user, err := s.GetUserByUsername(ctx, testUser.Username)
	assert.Nil(t, err)
	assert.Equal(t, testUser.Username, user.Username)

	// 测试空用户名
	user, err = s.GetUserByUsername(ctx, "")
	assert.NotNil(t, err)
	assert.Nil(t, user)

	// 测试不存在的用户名
	user, err = s.GetUserByUsername(ctx, "nonexistentuser")
	assert.NotNil(t, err)
	assert.Nil(t, user)
}
