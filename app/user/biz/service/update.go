package service

import (
	"context"
	"errors"

	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/model"
	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"
)

type UpdateService struct {
	ctx context.Context
}

// NewUpdateService new UpdateService
func NewUpdateService(ctx context.Context) *UpdateService {
	return &UpdateService{ctx: ctx}
}

// Run create note info
func (s *UpdateService) Run(req *user.UpdateReq) (resp *user.UpdateResp, err error) {
	// 将 string 转换为 int64

	// 检查用户是否存在
	existingUser, err := model.GetById(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, errors.New("user not found")
	}

	// 构建更新数据
	updates := &model.User{
		Username: req.UserName,
		Email:    req.Email,
		// ...
	}

	// 执行更新操作
	err = model.UpdateById(s.ctx, mysql.DB, req.UserId, updates)
	if err != nil {
		return nil, err
	}

	return &user.UpdateResp{}, nil
}
