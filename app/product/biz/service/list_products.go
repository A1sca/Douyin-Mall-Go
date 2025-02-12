package service

import (
	"context"

	"github.com/A1sca/Douyin-Mall-Go/app/api/biz/dal/mysql"
	product "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/product"
	"gorm.io/gorm"
)

type ListProductsService struct {
	ctx context.Context
}

// NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

func ListProducts(ctx context.Context, page, pageSize int) ([]*product.Product, int64, error) {
	// 计算偏移量
	offset := (page - 1) * pageSize

	// 初始化数据库连接
	var db *gorm.DB
	mysql.Init()

	// 定义变量来存储产品列表和总产品数
	var products []*product.Product
	var totalCount int64

	// 查询总产品数
	err := db.Model(&product.Product{}).Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	// 查询产品列表
	err = db.Limit(pageSize).Offset(offset).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return products, totalCount, nil
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {

	page := int(req.Page)
	pageSize := int(req.PageSize)

	ProductList, totalCount, err := ListProducts(s.ctx, page, pageSize)
	if err != nil {
		return nil, err
	}

	resp = &product.ListProductsResp{
		Products:   ProductList,
		TotalCount: int(totalCount),
	}

	return resp, nil
}
