package service

import (
	"context"

	"github.com/A1sca/Douyin-Mall-Go/app/api/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/product/biz/models"
	product "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/product"
	"gorm.io/gorm"
)

type ListProductsService struct {
	ctx context.Context
	db  *gorm.DB
}

// NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{
		ctx: ctx,
		db:  mysql.DB,
	}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {

	resp = &product.ListProductsResp{}

	// 构建数据库查询
	query := s.db

	// 如果请求中包含分类名称，添加分类过滤条件
	if req.CategoryName != "" {
		query = query.Where("category_name = ?", req.CategoryName)
	}

	// 计算偏移量
	offset := (int(req.Page - 1)) * int(req.PageSize)

	// 从数据库中查询产品列表
	var productModels []models.Product
	if err := query.Offset(int(offset)).Limit(int(req.PageSize)).Find(&productModels).Error; err != nil {
		return nil, err
	}

	// 将数据库中的产品列表转换为 ProtoBuf 定义的产品列表
	var products []*product.Product
	for _, productModel := range productModels {
		productInfo := &product.Product{
			Id:          uint32(productModel.ID),
			Name:        productModel.Name,
			Description: productModel.Description,
			Picture:     productModel.Picture,
			Price:       float32(productModel.Price),
			Categories:  productModel.Categories,
		}
		products = append(products, productInfo)
	}

	// 将产品列表添加到响应中
	resp.Products = products

	return resp, nil
}
