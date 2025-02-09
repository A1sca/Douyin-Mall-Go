package service

import (
	"context"
	"fmt"

	"github.com/A1sca/Douyin-Mall-Go/app/order/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/order/biz/model"

	"github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/cart"
	order "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/order"
)

type ListOrderService struct {
	ctx context.Context
}

// NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
// 根据用户id, 得到对应的订单列表
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// TODO: 检查用户是否存在
	userId := req.UserId
	orders, err := model.GetOrderListByUserId(s.ctx, mysql.DB, userId)
	fmt.Println(orders)
	if err != nil {
		return nil, err
	}
	rpcOrders := make([]*order.Order, len(orders))
	for i, modelOrder := range orders {
		rpcOrders[i] = fromModelOrder(&modelOrder)
	}
	return &order.ListOrderResp{
		Orders: rpcOrders,
	}, nil
}

func fromModelOrder(modelOrder *model.Order) *order.Order {
	rpcOrderItems := make([]*order.OrderItem, len(modelOrder.OrderItems))
	for i, modelOrderItem := range modelOrder.OrderItems {
		rpcOrderItems[i] = fromModelOrderItem(&modelOrderItem)
	}
	return &order.Order{
		OrderId:      modelOrder.OrderId,
		UserId:       modelOrder.UserId,
		UserCurrency: modelOrder.UserCurrency,
		Email:        modelOrder.Email,
		Address: &order.Address{
			StreetAddress: modelOrder.Address.StreetAddress,
			City:          modelOrder.Address.City,
			State:         modelOrder.Address.State,
			Country:       modelOrder.Address.Country,
			ZipCode:       modelOrder.Address.ZipCode,
		},
		OrderItems: rpcOrderItems,
	}
}

func fromModelOrderItem(modelOrderItem *model.OrderItem) *order.OrderItem {
	return &order.OrderItem{
		Cost: modelOrderItem.Cost,
		Item: &cart.CartItem{
			ProductId: modelOrderItem.ProductId,
			Quantity:  modelOrderItem.ProductNum,
		},
	}
}
