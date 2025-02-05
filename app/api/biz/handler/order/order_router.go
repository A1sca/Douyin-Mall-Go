package order

import (
	"context"

	"github.com/A1sca/Douyin-Mall-Go/app/api/biz/service"
	"github.com/A1sca/Douyin-Mall-Go/app/api/biz/utils"
	order "github.com/A1sca/Douyin-Mall-Go/app/api/hertz_gen/api/order"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CreateOrder .
// @router /v1/order/create [POST]
func CreateOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.CreateOrderReq

	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.CreateOrderResp{}
	resp, err = service.NewCreateOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateOrder .
// @router /v1/order/update [POST]
func UpdateOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.UpdateOrderReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.UpdateOrderResp{}
	resp, err = service.NewUpdateOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ListOrder .
// @router /v1/order/list/:user_id [GET]
func ListOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.ListOrderReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.ListOrderResp{}
	resp, err = service.NewListOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
