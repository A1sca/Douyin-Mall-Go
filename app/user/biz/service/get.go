package service

import (
	"context"

	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/model"
	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"
)

type GetService struct {
	ctx context.Context
} // NewGetService new GetService
func NewGetService(ctx context.Context) *GetService {
	return &GetService{ctx: ctx}
}

// Run create note info
func (s *GetService) Run(req *user.GetReq) (resp *user.GetResp, err error) {
	data, err := model.GetById(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return &user.GetResp{UserId: ""}, err
	}
	s.ctx = context.WithValue(s.ctx, "data", data)
	return &user.GetResp{UserId: req.UserId}, nil
}
