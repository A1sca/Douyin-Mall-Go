package order

import (
	"context"
	order "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func CreateOrder(ctx context.Context, req *order.CreateOrderReq, callOptions ...callopt.Option) (resp *order.CreateOrderResp, err error) {
	resp, err = defaultClient.CreateOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateOrder(ctx context.Context, req *order.UpdateOrderReq, callOptions ...callopt.Option) (resp *order.UpdateOrderResp, err error) {
	resp, err = defaultClient.UpdateOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CancelOrder(ctx context.Context, req *order.CancelOrderReq, callOptions ...callopt.Option) (resp *order.CancelOrderResp, err error) {
	resp, err = defaultClient.CancelOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CancelOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListOrder(ctx context.Context, req *order.ListOrderReq, callOptions ...callopt.Option) (resp *order.ListOrderResp, err error) {
	resp, err = defaultClient.ListOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
