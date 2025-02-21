package user

import (
	"context"
	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"

	"github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user/usermanageservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() usermanageservice.Client
	Service() string
	Get(ctx context.Context, Req *user.GetReq, callOptions ...callopt.Option) (r *user.GetResp, err error)
	Delete(ctx context.Context, Req *user.DeleteReq, callOptions ...callopt.Option) (r *user.DeleteResp, err error)
	Update(ctx context.Context, Req *user.UpdateReq, callOptions ...callopt.Option) (r *user.UpdateResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := usermanageservice.NewClient(dstService, opts...)
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
	kitexClient usermanageservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() usermanageservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Get(ctx context.Context, Req *user.GetReq, callOptions ...callopt.Option) (r *user.GetResp, err error) {
	return c.kitexClient.Get(ctx, Req, callOptions...)
}

func (c *clientImpl) Delete(ctx context.Context, Req *user.DeleteReq, callOptions ...callopt.Option) (r *user.DeleteResp, err error) {
	return c.kitexClient.Delete(ctx, Req, callOptions...)
}

func (c *clientImpl) Update(ctx context.Context, Req *user.UpdateReq, callOptions ...callopt.Option) (r *user.UpdateResp, err error) {
	return c.kitexClient.Update(ctx, Req, callOptions...)
}
