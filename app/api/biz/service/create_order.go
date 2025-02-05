package service

import (
	"context"

	order "github.com/A1sca/Douyin-Mall-Go/app/api/hertz_gen/api/order"
	"github.com/A1sca/Douyin-Mall-Go/app/api/rpc"
	rpc_order "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type CreateOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateOrderService(Context context.Context, RequestContext *app.RequestContext) *CreateOrderService {
	return &CreateOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateOrderService) Run(req *order.CreateOrderReq) (resp *order.CreateOrderResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	// fmt.Println("api 运行create order服务")
	// fmt.Printf("输出req.Address: %v\n", req.Address)
	// fmt.Printf("输出req.Address.Province: %v\n", req.Address.Province)
	// fmt.Printf("输出req.UserId: %v\n", req.UserId)
	// fmt.Printf("输出req.Orders: %v\n", req.Orders)

	orders := make([]*rpc_order.OrderItem, len(req.Orders))
	for i, item := range req.Orders {
		orders[i] = fromApiOrderItem(item)
	}
	rpc_resp, err := rpc.OrderService.CreateOrder(h.Context, &rpc_order.CreateOrderReq{
		UserId:  req.UserId,
		Address: fromApiAddress(req.Address),
		Orders:  orders,
	})
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderResp{
		OrderId: rpc_resp.OrderId,
	}, nil
}

func fromApiAddress(apiAddress *order.Address) *rpc_order.Address {
	return &rpc_order.Address{
		Phone:    apiAddress.Phone,
		Name:     apiAddress.Name,
		Province: apiAddress.Province,
		City:     apiAddress.City,
		District: apiAddress.District,
		Detail:   apiAddress.Detail,
	}
}

func fromApiOrderItem(apiOrderItem *order.OrderItem) *rpc_order.OrderItem {
	return &rpc_order.OrderItem{
		ProductId:  apiOrderItem.ProductId,
		ProductNum: apiOrderItem.ProductNum,
	}
}
