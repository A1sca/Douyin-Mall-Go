package user

import (
	"context"
	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"

	"github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() userservice.Client
	Service() string
	Register(ctx context.Context, Req *user.RegisterReq, callOptions ...callopt.Option) (r *user.RegisterResp, err error)
	Login(ctx context.Context, Req *user.LoginReq, callOptions ...callopt.Option) (r *user.LoginResp, err error)
	Logout(ctx context.Context, Req *user.LogoutReq, callOptions ...callopt.Option) (r *user.LogoutResp, err error)
	Get(ctx context.Context, Req *user.GetReq, callOptions ...callopt.Option) (r *user.GetResp, err error)
	Update(ctx context.Context, Req *user.UpdateReq, callOptions ...callopt.Option) (r *user.UpdateResp, err error)
	Delete(ctx context.Context, Req *user.DeleteReq, callOptions ...callopt.Option) (r *user.DeleteResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := userservice.NewClient(dstService, opts...)
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
	kitexClient userservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() userservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Register(ctx context.Context, Req *user.RegisterReq, callOptions ...callopt.Option) (r *user.RegisterResp, err error) {
	return c.kitexClient.Register(ctx, Req, callOptions...)
}

func (c *clientImpl) Login(ctx context.Context, Req *user.LoginReq, callOptions ...callopt.Option) (r *user.LoginResp, err error) {
	return c.kitexClient.Login(ctx, Req, callOptions...)
}

func (c *clientImpl) Logout(ctx context.Context, Req *user.LogoutReq, callOptions ...callopt.Option) (r *user.LogoutResp, err error) {
	return c.kitexClient.Logout(ctx, Req, callOptions...)
}

func (c *clientImpl) Get(ctx context.Context, Req *user.GetReq, callOptions ...callopt.Option) (r *user.GetResp, err error) {
	return c.kitexClient.Get(ctx, Req, callOptions...)
}

func (c *clientImpl) Update(ctx context.Context, Req *user.UpdateReq, callOptions ...callopt.Option) (r *user.UpdateResp, err error) {
	return c.kitexClient.Update(ctx, Req, callOptions...)
}

func (c *clientImpl) Delete(ctx context.Context, Req *user.DeleteReq, callOptions ...callopt.Option) (r *user.DeleteResp, err error) {
	return c.kitexClient.Delete(ctx, Req, callOptions...)
}
