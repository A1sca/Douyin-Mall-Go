package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/model"
	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"
)

type GetService struct {
	ctx context.Context
}

// NewGetService new GetService
func NewGetService(ctx context.Context) *GetService {
	return &GetService{ctx: ctx}
}

// Run create note info
func (s *GetService) Run(req *user.GetReq) (resp *user.GetResp, err error) {
	data, err := model.GetById(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, err
	}
	s.ctx = context.WithValue(s.ctx, "data", data)
	return &user.GetResp{}, nil
}

// GetUser 获取用户信息
func (s *GetService) GetUser(ctx context.Context, userID int64) (*model.User, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}

	// 将 int64 转换为字符串
	userIdStr := strconv.FormatInt(userID, 10)
	data, err := model.GetById(ctx, mysql.DB, userIdStr)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errors.New("user not found")
	}

	return data, nil
}

// GetUserByUsername 通过用户名获取用户信息
func (s *GetService) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	// 从数据库获取用户信息
	data, err := model.GetByName(ctx, mysql.DB, username) // 修改为正确的方法名
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errors.New("user not found")
	}

	return data, nil
}
