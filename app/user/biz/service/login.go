package service

import (
	"context"
	"errors"
	
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/utils"
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/model"
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/dal/mysql"
	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"
)

type LoginService struct {
	ctx context.Context
}

// NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("username or password cannot be empty")
	}

	// 获取用户信息
	user, err := model.GetByName(s.ctx, mysql.DB, req.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// 验证密码
	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("invalid password")
	}

	// 生成token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &user.LoginResp{
		UserId: user.ID,
		Token:  token,
	}, nil
}
