package service

import (
	"context"
	checkout "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/checkout"
)

type CheckoutService struct {
	ctx context.Context
}

// NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.

	return
}
