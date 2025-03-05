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

func TestUpdate_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateService(ctx)

	// 创建测试用户
	testUser := &model.User{
		Username: "updatetest",
		Email:    "updatetest@example.com",
		Phone:    "13800138007",
	}
	t.Logf("[测试准备] 创建测试用户: %+v", testUser)
	err := model.Create(ctx, mysql.DB, testUser)
	assert.Nil(t, err)

	// 测试正常更新
	req := &user.UpdateReq{
		UserId:   strconv.FormatUint(uint64(testUser.ID), 10),
		UserName: "updatedname",
		Email:    "updated@example.com",
	}
	t.Logf("[测试场景1] 正常更新用户信息: %+v", req)
	resp, err := s.Run(req)
	t.Logf("[测试结果1] 响应: %+v, 错误: %v", resp, err)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	// 测试更新不存在的用户
	req = &user.UpdateReq{
		UserId:   "99999",
		UserName: "nonexistent",
	}
	t.Logf("[测试场景2] 更新不存在的用户: %+v", req)
	resp, err = s.Run(req)
	t.Logf("[测试结果2] 响应: %+v, 错误: %v", resp, err)
	assert.NotNil(t, err)

	// 测试无效的更新数据
	req = &user.UpdateReq{
		UserId: strconv.FormatInt(int64(testUser.ID), 10),
	}
	t.Log("\n测试场景3: 更新数据为空")
	t.Logf("请求参数: %+v", req)
	resp, err = s.Run(req)
	t.Logf("响应: %+v, 错误: %v", resp, err)
	assert.Nil(t, err)
	t.Log("空数据更新测试通过")

	// 清理测试数据
	t.Log("\n清理测试数据...")
	err = model.DeleteByName(ctx, mysql.DB, "updatetest")
	assert.Nil(t, err)
	t.Log("测试数据清理完成")

	t.Log("=== 用户更新服务测试完成 ===")
}
