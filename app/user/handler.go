package main

import (
	"context"
	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}
type UserManageServiceImpl struct {}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Logout(ctx context.Context, req *user.LogoutReq) (resp *user.LogoutResp, err error) {
	resp, err = service.NewLogoutService(ctx).Run(req)

	return resp, err
}

// Get implements the UserManageServiceImpl interface.
func (s *UserManageServiceImpl) Get(ctx context.Context, req *user.GetReq) (resp *user.GetResp, err error) {
	resp, err = service.NewGetService(ctx).Run(req)

	return resp, err
}

// Delete implements the UserManageServiceImpl interface.
func (s *UserManageServiceImpl) Delete(ctx context.Context, req *user.DeleteReq) (resp *user.DeleteResp, err error) {
	resp, err = service.NewDeleteService(ctx).Run(req)

	return resp, err
}

// Update implements the UserManageServiceImpl interface.
func (s *UserManageServiceImpl) Update(ctx context.Context, req *user.UpdateReq) (resp *user.UpdateResp, err error) {
	resp, err = service.NewUpdateService(ctx).Run(req)

	return resp, err
}
