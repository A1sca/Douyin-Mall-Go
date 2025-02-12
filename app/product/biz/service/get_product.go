package service

import (
	"context"

	"github.com/A1sca/Douyin-Mall-Go/app/api/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/product/biz/models"
	product "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/product"
	"gorm.io/gorm"
)

type GetProductService struct {
	ctx context.Context
	db  *gorm.DB
}

// NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{
		ctx: ctx,
		db:  mysql.DB,
	}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp = &product.GetProductResp{}

	// 根据请求中的产品 ID 从数据库中查找产品
	var productModel models.Product
	if err := s.db.Where("id = ?", req.Id).First(&productModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 如果未找到产品，返回空响应
			return resp, nil
		}
		// 处理数据库查询错误
		return nil, err
	}

	// 将数据库中的产品信息转换为 ProtoBuf 定义的产品信息
	productInfo := &product.Product{
		Id:          uint32(productModel.ID),
		Name:        productModel.Name,
		Description: productModel.Description,
		Picture:     productModel.Picture,
		Price:       float32(productModel.Price),
		Categories:  productModel.Categories,
	}

	// 将产品信息添加到响应中
	resp.Product = productInfo

	return resp, nil
}
