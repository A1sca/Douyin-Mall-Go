package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/A1sca/Douyin-Mall-Go/app/api/biz/dal/mysql"
	product "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
}

// NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	resp = &product.SearchProductsResp{
		Results: make([]*product.Product, 0),
	}

	query := req.Query
	if query == "" {
		return resp, nil
	}
	query = strings.TrimSpace(strings.ToLower(query))

	products, err := mysql.SearchProductsByKeyword(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("未查到商品: %w", err)
	}

	// 将搜索结果添加到响应中
	for _, p := range products {
		resp.Results = append(resp.Results, &product.Product{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
			Categories:  p.Categories,
		})
	}
	return resp, nil
}
