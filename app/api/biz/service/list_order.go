package service

import (
	"context"

	api_cart "github.com/A1sca/Douyin-Mall-Go/app/api/hertz_gen/api/cart"
	api_order "github.com/A1sca/Douyin-Mall-Go/app/api/hertz_gen/api/order"
	"github.com/A1sca/Douyin-Mall-Go/app/api/rpc"
	rpc_order "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type ListOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListOrderService(Context context.Context, RequestContext *app.RequestContext) *ListOrderService {
	return &ListOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *ListOrderService) Run(req *api_order.ListOrderReq) (resp *api_order.ListOrderResp, err error) {
	userId := req.UserId
	// fmt.Println("用户id: ", userId)
	rpc_resp, err := rpc.OrderService.ListOrder(h.Context, &rpc_order.ListOrderReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	apiOrders := make([]*api_order.Order, len(rpc_resp.Orders))
	for i, rpcOrder := range rpc_resp.Orders {
		apiOrders[i] = fromRpcOrder(rpcOrder)
	}
	return &api_order.ListOrderResp{
		Orders: apiOrders,
	}, nil
}

func fromRpcOrder(rpcOrder *rpc_order.Order) *api_order.Order {
	apiOrderItems := make([]*api_order.OrderItem, len(rpcOrder.OrderItems))
	for i, rpcOrderItem := range rpcOrder.OrderItems {
		apiOrderItems[i] = fromRpcOrderItem(rpcOrderItem)
	}
	return &api_order.Order{
		OrderId:      rpcOrder.OrderId,
		UserId:       rpcOrder.UserId,
		UserCurrency: rpcOrder.UserCurrency,
		Address: &api_order.Address{
			StreetAddress: rpcOrder.Address.StreetAddress,
			City:          rpcOrder.Address.City,
			State:         rpcOrder.Address.State,
			Country:       rpcOrder.Address.Country,
			ZipCode:       rpcOrder.Address.ZipCode,
		},
		Email:      rpcOrder.Email,
		OrderItems: apiOrderItems,
	}
}

func fromRpcOrderItem(rpcItem *rpc_order.OrderItem) *api_order.OrderItem {
	return &api_order.OrderItem{
		Cost: rpcItem.Cost,
		Item: &api_cart.CartItem{
			ProductId: rpcItem.Item.ProductId,
			Quantity:  rpcItem.Item.Quantity,
		},
	}
}
