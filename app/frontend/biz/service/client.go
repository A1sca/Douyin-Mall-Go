package service

import (
	"time"

	"github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
)

var (
	userClient     userservice.Client
	sessionTimeout = 24 * time.Hour
)

// InitRPCClients 初始化所有 RPC 客户端
func InitRPCClients() {
	var err error
	userClient, err = userservice.NewClient("user", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		panic(err)
	}
}
