package service

import (
	"context"
	"fmt"

	"github.com/A1sca/Douyin-Mall-Go/app/api/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/product/biz/models"
	product "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/product"
	"gorm.io/gorm"
)

type SearchProductsService struct {
	ctx context.Context
	db  *gorm.DB
}

// NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{
		ctx: ctx,
		db:  mysql.DB,
	}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	resp = &product.SearchProductsResp{}
	// 获取查询字符串并进行处理
	query := req.Query
	if query == "" {
		return resp, nil // 如果查询字符串为空，直接返回空结果
	}

	// 假设 mysql.SearchProducts 是一个在数据库中搜索产品的函数
	// 它接收上下文和查询字符串作为参数，返回产品列表和错误
	products, err := models.SearchProducts(s.ctx, s.db, query)
	if err != nil {
		return nil, fmt.Errorf("failed to search products: %w", err)
	}

	// 将搜索到的产品转换为 ProtoBuf 定义的产品
	var protoProducts []*product.Product
	for _, p := range products {
		protoProduct := &product.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       float32(p.Price),
			Categories:  p.Categories,
		}
		protoProducts = append(protoProducts, protoProduct)
	}

	// 将转换后的产品添加到响应中
	resp.Results = protoProducts

	return resp, nil
}
