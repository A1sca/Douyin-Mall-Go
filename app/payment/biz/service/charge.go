package service

import (
	"context"
	payment "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/payment"
)

type ChargeService struct {
	ctx context.Context
}

// NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// Finish your business logic.

	return
}
