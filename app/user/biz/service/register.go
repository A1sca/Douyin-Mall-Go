package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/model"
	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
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
	// Finish your business logic.
	if req.Email == "" || req.Password == "" || req.PasswordConfirm == "" {
		return nil, errors.New("email or password is empty")
	}
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("password not match")
	}
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &model.User{
		Username:       "default user name",
		Phone:          "123",
		Gender:         1,
		Age:            1,
		Avatar:         "https://picsum.photos/id/237/200/300",
		Email:          req.Email,
		Balance:        0,
		Desc:           "there is nothing",
		Addresses:      make([]model.Address, 0),
		PasswordHashed: string(passwordHashed),
	}
	err = model.Create(s.ctx, mysql.DB, newUser)
	if err != nil {
		return nil, err
	}
	return &user.RegisterResp{UserId: strconv.FormatUint((uint64(newUser.ID)), 10)}, nil
}
