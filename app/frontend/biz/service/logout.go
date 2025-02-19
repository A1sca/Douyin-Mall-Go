package service

import (
	"context"

	"github.com/A1sca/Douyin-Mall-Go/app/frontend/biz/dal/redis"
	auth "github.com/A1sca/Douyin-Mall-Go/app/frontend/hertz_gen/frontend/auth"
	common "github.com/A1sca/Douyin-Mall-Go/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol"
)

type LogoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLogoutService(Context context.Context, RequestContext *app.RequestContext) *LogoutService {
	return &LogoutService{RequestContext: RequestContext, Context: Context}
}

func (h *LogoutService) Run(req *auth.LogoutReq) (resp *common.Empty, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "logout request processed")
	}()

	// 获取当前会话ID
	sessionID := string(h.RequestContext.Cookie("session_id"))
	if sessionID != "" {
		// 从 Redis 中删除会话
		err := redis.RedisClient.Del(h.Context, sessionID).Err()
		if err != nil {
			hlog.CtxWarnf(h.Context, "failed to delete session: %v", err)
		}
	}

	// 删除 cookie
	h.RequestContext.SetCookie(
		"session_id",
		"",
		-1, // 立即过期
		"/",
		"",
		protocol.CookieSameSiteDisabled, // 使用正确的 sameSite 值
		false,
		true,
	)

	return &common.Empty{}, nil
}
