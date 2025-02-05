package rpc

import (
	"sync"

	"github.com/A1sca/Douyin-Mall-Go/app/api/conf"
	"github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/auth/authservice"
	"github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/order/orderservice"
	"github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/payment/paymentservice"
	productservice "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	AuthService     authservice.Client
	UserService     userservice.Client
	ProductService  productservice.Client
	CartService     cartservice.Client
	OrderService    orderservice.Client
	CheckoutService checkoutservice.Client
	PaymentService  paymentservice.Client
	once            sync.Once
)

func InitClient() {
	once.Do(func() {
		initAuthServiceClient()
		initUserServiceClient()
		initProductServiceClient()
		initCartServiceClient()
		initOrderServiceClient()
		initCheckoutServiceClient()
		initPaymentServiceClient()
	})
}

func initAuthServiceClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	if err != nil {
		hlog.Fatal(err)
	}
	AuthService, err = authservice.NewClient("auth", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}
}

func initUserServiceClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	if err != nil {
		hlog.Fatal(err)
	}
	UserService, err = userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}
}

func initProductServiceClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	if err != nil {
		hlog.Fatal(err)
	}
	ProductService, err = productservice.NewClient("product", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}
}

func initCartServiceClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	if err != nil {
		hlog.Fatal(err)
	}
	CartService, err = cartservice.NewClient("cart", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}
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

func initCheckoutServiceClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	if err != nil {
		hlog.Fatal(err)
	}
	CheckoutService, err = checkoutservice.NewClient("checkout", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}
}

func initPaymentServiceClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	if err != nil {
		hlog.Fatal(err)
	}
	PaymentService, err = paymentservice.NewClient("payment", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}
}
