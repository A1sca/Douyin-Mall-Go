package service

import (
	"context"

	"github.com/A1sca/Douyin-Mall-Go/app/order/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/order/biz/model"
	order "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/order"
	"github.com/google/uuid"
)

type PlaceOrderService struct {
	ctx context.Context
}

// NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// 各种检查 ...
	// 检查用户id
	// TODO: 检查用户余额(需要调用其他服务)
	// 生成id
	new_order_id := uuid.New().String()
	// 把order.CreateOrderReq的参数转为model的参数
	new_order_address := fromProtoAddress(req.Address)
	new_order_items := make([]model.OrderItem, len(req.OrderItems))
	for i, item := range req.OrderItems {
		new_order_items[i] = fromProtoOrderItem(item)
	}
	// 构造model.Order
	new_order := &model.Order{
		OrderId:      new_order_id,
		UserId:       req.UserId,
		UserCurrency: req.UserCurrency,
		Email:        req.Email,
		Address:      new_order_address,
		OrderItems:   new_order_items,
	}
	err = model.Create(s.ctx, mysql.DB, new_order)
	if err != nil {
		return nil, err
	}
	return &order.PlaceOrderResp{
		Order: &order.OrderResult{
			OrderId: new_order_id,
		},
	}, nil
}

func fromProtoAddress(protoAddress *order.Address) model.Address {
	return model.Address{
		StreetAddress: protoAddress.StreetAddress,
		City:          protoAddress.City,
		State:         protoAddress.State,
		Country:       protoAddress.Country,
		ZipCode:       protoAddress.ZipCode,
	}
}

func fromProtoOrderItem(protoOrderItem *order.OrderItem) model.OrderItem {
	return model.OrderItem{
		ProductId:  protoOrderItem.Item.ProductId,
		ProductNum: protoOrderItem.Item.Quantity,
		Cost:       protoOrderItem.Cost,
	}
}
