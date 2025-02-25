package service

import (
	"context"
	"fmt"

	"github.com/A1sca/Douyin-Mall-Go/app/api/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/product/biz/models"
	product "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/product/kitex_gen"
	"gorm.io/gorm"
)

type CreateProductService struct {
	ctx context.Context
	db  *gorm.DB
}

// NewGetProductService new GetProductService
func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{
		ctx: ctx,
		db:  mysql.DB,
	}
}

func (s *CreateProductService) Run(req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	resp = &product.CreateProductResp{}
	productModel := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  req.Categories,
	}
	if err := s.db.Create(productModel).Error; err != nil {
		return nil, fmt.Errorf("未能生成: %w", err)
	}

	productInfo := &product.Product{
		Id:          uint32(productModel.ID),
		Name:        productModel.Name,
		Description: productModel.Description,
		Picture:     productModel.Picture,
		Price:       float32(productModel.Price),
		Categories:  productModel.Categories,
	}
	resp.Product = productInfo

	return resp, nil
}
