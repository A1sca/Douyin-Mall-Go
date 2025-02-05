package rpc

import (
	"sync"

	"github.com/A1sca/Douyin-Mall-Go/app/api/conf"
	"github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	OrderService orderservice.Client
	once         sync.Once
)

func InitClient() {
	once.Do(func() {
		initOrderServiceClient()
	})
}

func initOrderServiceClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	if err != nil {
		hlog.Fatal(err)
	}
	OrderService, err = orderservice.NewClient("order", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}
}
