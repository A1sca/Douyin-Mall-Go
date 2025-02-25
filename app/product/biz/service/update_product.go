package service

import (
	"context"

	"github.com/A1sca/Douyin-Mall-Go/app/api/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/product/biz/models"
	product "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/product/kitex_gen"
	"gorm.io/gorm"
)

type UpdateProductService struct {
	ctx context.Context
	db  *gorm.DB
}

// NewGetProductService new GetProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
	return &UpdateProductService{
		ctx: ctx,
		db:  mysql.DB,
	}
}

func (s *UpdateProductService) Run(req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	resp = new(product.UpdateProductResp)

	var productModel models.Product
	if err := s.db.Where("id = ?", req.Id).First(&productModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 如果未找到产品，返回错误响应
			resp.success = false
			resp.message = "产品不存在"
			return resp, nil
		}
		return nil, err
	}
	// 处理数据库查询错误
	if req.Name != "" {
		productModel.Name = req.Name
	}
	if req.Description != "" {
		productModel.Description = req.Description
	}
	if req.Picture != "" {
		productModel.Picture = req.Picture
	}
	if req.Price != 0 {
		productModel.Price = float32(req.Price)
	}
	if req.Categories != nil {
		productModel.Categories = req.Categories
	}

	// 保存更新后的产品信息到数据库
	if err := s.db.Save(&productModel).Error; err != nil {
		return nil, err
	}

	// 设置响应状态为成功
	resp.Success = true
	resp.Message = "产品更新成功"
	return resp, nil
}
