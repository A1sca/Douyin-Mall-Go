package service

import (
	"context"
	"fmt"

	order "github.com/A1sca/Douyin-Mall-Go/app/api/hertz_gen/api/order"
	"github.com/A1sca/Douyin-Mall-Go/app/api/rpc"
	rpc_cart "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/cart"
	rpc_order "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/order"

	"github.com/cloudwego/hertz/pkg/app"
)

type PlaceOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPlaceOrderService(Context context.Context, RequestContext *app.RequestContext) *PlaceOrderService {
	return &PlaceOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	fmt.Println("运行Place Order")
	fmt.Println(req.UserId)
	fmt.Println(req.UserCurrency)
	fmt.Println(req.Address)
	fmt.Println(req.Email)
	fmt.Println(req.OrderItems)
	orderItems := make([]*rpc_order.OrderItem, len(req.OrderItems))
	for i, item := range req.OrderItems {
		orderItems[i] = fromApiOrderItem(item)
	}
	r, err := rpc.OrderService.PlaceOrder(h.Context, &rpc_order.PlaceOrderReq{
		UserId:       req.UserId,
		UserCurrency: req.UserCurrency,
		Address:      fromApiAddress(req.Address),
		Email:        req.Email,
		OrderItems:   orderItems,
	})
	if err != nil {
		return nil, err
	}
	return &order.PlaceOrderResp{
		Order: &order.OrderResult{
			OrderId: r.Order.OrderId,
		},
	}, err
}

func fromApiAddress(apiAddress *order.Address) *rpc_order.Address {
	return &rpc_order.Address{
		StreetAddress: apiAddress.StreetAddress,
		City:          apiAddress.City,
		State:         apiAddress.State,
		Country:       apiAddress.Country,
		ZipCode:       apiAddress.ZipCode,
	}
}

func fromApiOrderItem(apiOrderItem *order.OrderItem) *rpc_order.OrderItem {
	return &rpc_order.OrderItem{
		Item: &rpc_cart.CartItem{
			ProductId: apiOrderItem.Item.ProductId,
			Quantity:  apiOrderItem.Item.Quantity,
		},
		Cost: apiOrderItem.Cost,
	}
}
