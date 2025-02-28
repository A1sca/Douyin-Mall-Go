package middleware

import (
	"context"
	"net/http"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/stretchr/testify/assert"
)

func TestJWTAuthMiddleware(t *testing.T) {
	middleware := JWTAuthMiddleware()
	
	tests := []struct {
		name          string
		token         string
		expectedCode  int
		expectedAbort bool
	}{
		{
			name:          "无token",
			token:         "",
			expectedCode:  http.StatusUnauthorized,
			expectedAbort: true,
		},
		{
			name:          "token格式错误",
			token:         "invalid_token",
			expectedCode:  http.StatusUnauthorized,
			expectedAbort: true,
		},
		{
			name:          "有效token",
			token:         "Bearer valid_token",
			expectedCode:  http.StatusOK,
			expectedAbort: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建测试上下文
			ctx := context.Background()
			c := app.RequestContext{
				Request:  protocol.Request{},
				Response: protocol.Response{},
			}
			
			if tt.token != "" {
				c.Request.Header.Set("Authorization", tt.token)
			}
			
			// 执行中间件
			middleware(ctx, &c)
			
			// 验证结果
			assert.Equal(t, tt.expectedCode, c.Response.StatusCode())
			if tt.expectedAbort {
				// 验证是否中止了请求处理
				assert.True(t, c.IsAborted())
			}
		})
	}
} 