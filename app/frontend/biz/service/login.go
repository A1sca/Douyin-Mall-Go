package service

import (
	"context"
	"errors"

	"github.com/A1sca/Douyin-Mall-Go/app/frontend/biz/dal/redis"
	"github.com/A1sca/Douyin-Mall-Go/app/frontend/biz/utils"
	auth "github.com/A1sca/Douyin-Mall-Go/app/frontend/hertz_gen/frontend/auth"
	common "github.com/A1sca/Douyin-Mall-Go/app/frontend/hertz_gen/frontend/common"
	user "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (resp *common.Empty, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	// 参数验证
	if err := h.validateLoginRequest(req); err != nil {
		return nil, err
	}

	// 检查是否已经登录
	if h.isAlreadyLoggedIn() {
		return nil, errors.New("already logged in")
	}

	// 调用用户服务进行登录验证
	userResp, err := h.doLogin(req)
	if err != nil {
		return nil, err
	}

	// 设置登录会话
	if err := h.setLoginSession(userResp.UserId); err != nil {
		return nil, err
	}

	return &common.Empty{}, nil
}

// validateLoginRequest 验证登录请求参数
func (h *LoginService) validateLoginRequest(req *auth.LoginReq) error {
	if req.Email == "" {
		return errors.New("email cannot be empty")
	}
	if req.Password == "" {
		return errors.New("password cannot be empty")
	}
	return nil
}

// isAlreadyLoggedIn 检查是否已经登录
func (h *LoginService) isAlreadyLoggedIn() bool {
	sessionID := h.RequestContext.Cookie("session_id")
	if len(sessionID) == 0 {
		return false
	}

	userID, err := utils.GetUserIDFromSession(h.Context, string(sessionID))
	return err == nil && userID != ""
}

// doLogin 执行登录操作
func (h *LoginService) doLogin(req *auth.LoginReq) (*user.LoginResp, error) {
	userReq := &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	}

	userResp, err := userClient.Login(h.Context, userReq)
	if err != nil {
		hlog.CtxErrorf(h.Context, "login failed: %v", err)
		return nil, errors.New("login failed")
	}

	return userResp, nil
}

// setLoginSession 设置登录会话
func (h *LoginService) setLoginSession(userID string) error {
	sessionID := utils.GenerateSessionID()

	// 在 Redis 中存储会话信息
	err := redis.RedisClient.Set(h.Context, sessionID, userID, sessionTimeout).Err()
	if err != nil {
		hlog.CtxErrorf(h.Context, "failed to set session: %v", err)
		return errors.New("failed to set session")
	}

	// 设置 cookie
	h.RequestContext.SetCookie(
		"session_id",
		sessionID,
		int(sessionTimeout.Seconds()),
		"/",
		"",
		protocol.CookieSameSiteDisabled,
		false, // secure
		true,  // httpOnly
	)

	return nil
}

// 其他辅助函数...
