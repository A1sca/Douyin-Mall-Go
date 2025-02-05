package service

import (
	"context"

	"github.com/A1sca/Douyin-Mall-Go/app/order/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/order/biz/model"
	order "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/order"
	"github.com/google/uuid"
)

type CreateOrderService struct {
	ctx context.Context
}

// NewCreateOrderService new CreateOrderService
func NewCreateOrderService(ctx context.Context) *CreateOrderService {
	return &CreateOrderService{ctx: ctx}
}

// Run create note info
func (s *CreateOrderService) Run(req *order.CreateOrderReq) (resp *order.CreateOrderResp, err error) {
	// 各种检查 ...
	// 检查用户id
	// TODO: 检查用户余额(需要调用其他服务)
	// 生成id
	new_order_id := uuid.New().String()
	// 把order.CreateOrderReq的参数转为model的参数
	new_order_address := fromProtoAddress(req.Address)
	new_order_items := make([]model.OrderItem, len(req.Orders))
	for i, item := range req.Orders {
		new_order_items[i] = fromProtoOrderItem(item)
	}
	// 构造model.Order
	new_order := &model.Order{
		OrderId:      new_order_id,
		UserId:       req.UserId,
		OrderAddress: new_order_address,
		OrderItems:   new_order_items,
	}
	err = model.Create(s.ctx, mysql.DB, new_order)
	if err != nil {
		return nil, err
	}
	return &order.CreateOrderResp{OrderId: new_order.OrderId}, nil
}

func fromProtoAddress(protoAddress *order.Address) model.Address {
	return model.Address{
		Phone:    protoAddress.Phone,
		Name:     protoAddress.Name,
		Province: protoAddress.Province,
		City:     protoAddress.City,
		District: protoAddress.District,
		Detail:   protoAddress.Detail,
	}
}

func fromProtoOrderItem(protoOrderItem *order.OrderItem) model.OrderItem {
	return model.OrderItem{
		ProductId:  protoOrderItem.ProductId,
		ProductNum: protoOrderItem.ProductNum,
	}
}
