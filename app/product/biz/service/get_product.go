package service

import (
	"context"

	"github.com/A1sca/Douyin-Mall-Go/app/api/biz/dal/mysql"
	product "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/product"
	"gorm.io/gorm"
)

type GetProductService struct {
	ctx context.Context
	db  *gorm.DB
}

// NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	db := product.getDB()
	return &GetProductService{
		ctx: ctx,
		db:  db,
	}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	productID := req.Id

	// 调用数据访问层方法获取产品信息
	productInfo, err := mysql.GetProductByID(s.ctx, productID)
	if err != nil {
		// 处理错误，返回错误响应
		return nil, err
	}

	// 构建响应
	resp = &product.GetProductResp{
		Product: productInfo,
	}

	return resp, nil
}
