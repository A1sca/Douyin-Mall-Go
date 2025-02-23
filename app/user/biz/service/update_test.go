package service

import (
	"context"
	"testing"

	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/model"
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/dal/mysql"
	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"
	"github.com/stretchr/testify/assert"
)

func TestUpdate_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateService(ctx)

	// 创建测试用户
	testUser := &model.User{
		Username: "updatetest",
		Email:    "updatetest@example.com",
		Phone:    "13800138007",
	}
	err := model.Create(ctx, mysql.DB, testUser)
	assert.Nil(t, err)

	// 测试正常更新
	req := &user.UpdateReq{
		UserId:   testUser.ID,
		UserName: "updatedname",
		Email:    "updated@example.com",
	}
	resp, err := s.Run(req)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	// 测试更新不存在的用户
	req = &user.UpdateReq{
		UserId:   99999,
		UserName: "nonexistent",
	}
	resp, err = s.Run(req)
	assert.NotNil(t, err)

	// 测试无效的更新数据
	req = &user.UpdateReq{
		UserId: testUser.ID,
	}
	resp, err = s.Run(req)
	assert.Nil(t, err)
}
