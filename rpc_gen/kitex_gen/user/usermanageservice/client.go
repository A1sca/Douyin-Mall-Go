// Code generated by Kitex v0.9.1. DO NOT EDIT.

package usermanageservice

import (
	"context"
	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Get(ctx context.Context, Req *user.GetReq, callOptions ...callopt.Option) (r *user.GetResp, err error)
	Delete(ctx context.Context, Req *user.DeleteReq, callOptions ...callopt.Option) (r *user.DeleteResp, err error)
	Update(ctx context.Context, Req *user.UpdateReq, callOptions ...callopt.Option) (r *user.UpdateResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserManageServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserManageServiceClient struct {
	*kClient
}

func (p *kUserManageServiceClient) Get(ctx context.Context, Req *user.GetReq, callOptions ...callopt.Option) (r *user.GetResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Get(ctx, Req)
}

func (p *kUserManageServiceClient) Delete(ctx context.Context, Req *user.DeleteReq, callOptions ...callopt.Option) (r *user.DeleteResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Delete(ctx, Req)
}

func (p *kUserManageServiceClient) Update(ctx context.Context, Req *user.UpdateReq, callOptions ...callopt.Option) (r *user.UpdateResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Update(ctx, Req)
}
