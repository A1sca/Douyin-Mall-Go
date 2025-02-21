package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/model"
	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"
)

type DeleteService struct {
	ctx context.Context
}

// NewDeleteService new DeleteService
func NewDeleteService(ctx context.Context) *DeleteService {
	return &DeleteService{ctx: ctx}
}

// Run create note info
func (s *DeleteService) Run(req *user.DeleteReq) (resp *user.DeleteResp, err error) {
	err = model.DeleteById(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, err
	}
	return &user.DeleteResp{}, nil
}

// DeleteUser 删除用户
func (s *DeleteService) DeleteUser(ctx context.Context, userID int64) error {
	if userID <= 0 {
		return errors.New("invalid user id")
	}

	// 将 int64 转换为字符串
	userIdStr := strconv.FormatInt(userID, 10)

	// 检查用户是否存在
	data, err := model.GetById(ctx, mysql.DB, userIdStr)
	if err != nil {
		return err
	}
	if data == nil {
		return errors.New("user not found")
	}

	// 执行删除操作
	err = model.DeleteById(ctx, mysql.DB, userIdStr)
	if err != nil {
		return err
	}

	return nil
}
