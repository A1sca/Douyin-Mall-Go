package user

import (
	"context"
	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Register(ctx context.Context, req *user.RegisterReq, callOptions ...callopt.Option) (resp *user.RegisterResp, err error) {
	resp, err = defaultClient.Register(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Register call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Login(ctx context.Context, req *user.LoginReq, callOptions ...callopt.Option) (resp *user.LoginResp, err error) {
	resp, err = defaultClient.Login(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Login call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Logout(ctx context.Context, req *user.LogoutReq, callOptions ...callopt.Option) (resp *user.LogoutResp, err error) {
	resp, err = defaultClient.Logout(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Logout call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Get(ctx context.Context, req *user.GetReq, callOptions ...callopt.Option) (resp *user.GetResp, err error) {
	resp, err = defaultClient.Get(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Get call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Update(ctx context.Context, req *user.UpdateReq, callOptions ...callopt.Option) (resp *user.UpdateResp, err error) {
	resp, err = defaultClient.Update(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Update call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Delete(ctx context.Context, req *user.DeleteReq, callOptions ...callopt.Option) (resp *user.DeleteResp, err error) {
	resp, err = defaultClient.Delete(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Delete call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
