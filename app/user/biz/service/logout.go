package service

import (
	"context"

	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"
)

type LogoutService struct {
	ctx context.Context
}

// NewLoginService new LoginService
func NewLogoutService(ctx context.Context) *LogoutService {
	return &LogoutService{ctx: ctx}
}

// Run create note info
func (s *LogoutService) Run(req *user.LogoutReq) (resp *user.LogoutResp, err error) {
	resp = &user.LogoutResp{
		UserId: req.UserId,
	}
	return resp, nil
}
