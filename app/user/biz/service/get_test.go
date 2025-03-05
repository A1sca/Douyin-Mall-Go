package service

import (
	"context"
	"strconv"
	"testing"

	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/model"
	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"
)

func TestGet_Run(t *testing.T) {
	t.Log("开始测试用户获取服务")
	ctx := context.Background()
	s := NewGetService(ctx)

	// 创建测试用户
	testUser := &model.User{
		Username: "gettest",
		Email:    "gettest@example.com",
		Phone:    "13800138003",
	}
	err := model.Create(ctx, mysql.DB, testUser)
	if err != nil {
		t.Fatalf("创建测试用户失败: %v", err)
	}
	t.Logf("成功创建测试用户: ID=%d, Username=%s", testUser.ID, testUser.Username)

	// 测试正常获取用户
	t.Log("测试场景1: 正常获取用户信息")
	req := &user.GetReq{UserId: strconv.FormatUint(uint64(testUser.ID), 10)}
	t.Logf("请求参数: %+v", req)
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("获取用户信息失败: %v", err)
	} else {
		t.Logf("获取用户信息成功: %+v", resp)
	}

	// 测试获取不存在的用户
	t.Log("\n测试场景2: 获取不存在的用户")
	req = &user.GetReq{UserId: "99999"}
	t.Logf("请求参数: %+v", req)
	resp, err = s.Run(req)
	if err != nil {
		t.Logf("预期的错误: %v", err)
	} else {
		t.Error("期望获取不存在用户返回错误，但是成功了")
	}

	t.Log("用户获取服务测试完成")
}
