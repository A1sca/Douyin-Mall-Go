package auth

import (
	"context"

	"github.com/A1sca/Douyin-Mall-Go/app/frontend/biz/service"
	"github.com/A1sca/Douyin-Mall-Go/app/frontend/biz/utils"
	auth "github.com/A1sca/Douyin-Mall-Go/app/frontend/hertz_gen/frontend/auth"
	common "github.com/A1sca/Douyin-Mall-Go/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Login .
// @router /auth/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.LoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &common.Empty{}
	resp, err = service.NewLoginService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
