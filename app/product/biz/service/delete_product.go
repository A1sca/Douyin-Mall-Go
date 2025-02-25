package service

import (
	"context"

	"github.com/A1sca/Douyin-Mall-Go/app/api/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/product/biz/models"
	product "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/product/kitex_gen"
	"gorm.io/gorm"
)

type DeleteProductService struct {
	ctx context.Context
	db  *gorm.DB
}

// NewGetProductService new GetProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{
		ctx: ctx,
		db:  mysql.DB,
	}
}

func (s *DeleteProductService) Run(req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	resp = &product.DeleteProductResp{}

	if err := s.db.Where("id = ?", req.Id).Delete(&models.Product{}).Error; err != nil {
		return nil, err
	}
	resp.Success = true
	return resp, nil
}
