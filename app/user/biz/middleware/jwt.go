package middleware

import (
	"context"
	"net/http"
	"strings"

	auth "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/auth"
	"github.com/A1sca/Douyin-Mall-Go/rpc_gen/rpc/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// JWTAuthMiddleware 用于验证JWT令牌的中间件
func JWTAuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从请求头中获取token
		authorization := string(c.GetHeader("Authorization"))
		if authorization == "" {
			c.JSON(http.StatusUnauthorized, utils.H{
				"code": 401,
				"msg":  "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		// 解析token
		parts := strings.SplitN(authorization, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, utils.H{
				"code": 401,
				"msg":  "token格式错误",
			})
			c.Abort()
			return
		}

		// 调用auth服务验证token
		resp, err := auth.VerifyTokenByRPC(ctx, &auth.VerifyTokenReq{Token: parts[1]})
		if err != nil || !resp.Res {
			c.JSON(http.StatusUnauthorized, utils.H{
				"code": 401,
				"msg":  "无效的token",
			})
			c.Abort()
			return
		}

		// 继续处理请求
		c.Next(ctx)
	}
}