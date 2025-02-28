package service

import (
	"context"
	auth "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/auth"
	"github.com/A1sca/Douyin-Mall-Go/app/auth/biz/utils"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
}

// NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// 验证token
	_, err = utils.ParseToken(req.Token)
	if err != nil {
		return &auth.VerifyResp{Res: false}, nil
	}
	return &auth.VerifyResp{Res: true}, nil
}
