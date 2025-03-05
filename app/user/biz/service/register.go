package service

import (
	"context"
	"errors"
	"log"
	"strconv"

	jwtutils "github.com/A1sca/Douyin-Mall-Go/app/auth/biz/utils"
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/model"
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/utils"
	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"
)

type RegisterService struct {
	ctx context.Context
}

// NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	defer func() {
		log.Printf("[RegisterService] req = %+v", req)
		log.Printf("[RegisterService] resp = %+v, err = %v", resp, err)
	}()
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("username or password cannot be empty")
	}

	// 检查用户名是否已存在
	existingUser, err := model.GetByName(s.ctx, mysql.DB, req.Username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	// 创建新用户
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		Username:       req.Username,
		PasswordHashed: hashedPassword,
		Email:          req.Email,
	}

	err = model.Create(s.ctx, mysql.DB, newUser)
	if err != nil {
		return nil, err
	}

	// 生成token
	token, err := jwtutils.GenerateToken(strconv.FormatInt(int64(newUser.ID), 10), newUser.Username)
	if err != nil {
		return nil, err
	}

	return &user.RegisterResp{
		UserId: strconv.FormatInt(int64(newUser.ID), 10),
		Token:  token,
	}, nil
}
