package order

import (
	"context"
	order "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/order"

	"github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() orderservice.Client
	Service() string
	CreateOrder(ctx context.Context, Req *order.CreateOrderReq, callOptions ...callopt.Option) (r *order.CreateOrderResp, err error)
	UpdateOrder(ctx context.Context, Req *order.UpdateOrderReq, callOptions ...callopt.Option) (r *order.UpdateOrderResp, err error)
	CancelOrder(ctx context.Context, Req *order.CancelOrderReq, callOptions ...callopt.Option) (r *order.CancelOrderResp, err error)
	ListOrder(ctx context.Context, Req *order.ListOrderReq, callOptions ...callopt.Option) (r *order.ListOrderResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := orderservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient orderservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() orderservice.Client {
	return c.kitexClient
}

func (c *clientImpl) CreateOrder(ctx context.Context, Req *order.CreateOrderReq, callOptions ...callopt.Option) (r *order.CreateOrderResp, err error) {
	return c.kitexClient.CreateOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateOrder(ctx context.Context, Req *order.UpdateOrderReq, callOptions ...callopt.Option) (r *order.UpdateOrderResp, err error) {
	return c.kitexClient.UpdateOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) CancelOrder(ctx context.Context, Req *order.CancelOrderReq, callOptions ...callopt.Option) (r *order.CancelOrderResp, err error) {
	return c.kitexClient.CancelOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) ListOrder(ctx context.Context, Req *order.ListOrderReq, callOptions ...callopt.Option) (r *order.ListOrderResp, err error) {
	return c.kitexClient.ListOrder(ctx, Req, callOptions...)
}
