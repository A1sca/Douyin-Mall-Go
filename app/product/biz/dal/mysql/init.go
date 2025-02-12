package mysql

import (
	"context"
	"fmt"

	"github.com/A1sca/Douyin-Mall-Go/app/product/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	product "github.com/A1sca/Douyin-Mall-Go/rpc_gen/kitex_gen/product"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
}

type Product struct {
	ID          uint32 `gorm:"primaryKey"`
	Name        string
	Description string
	Picture     string
	Price       float32
	Categories  []string `gorm:"serializer:json"`
}

func SearchProductsByKeyword(ctx context.Context, keyword string) ([]*Product, error) {
	var products []*Product
	err := DB.WithContext(ctx).Where("LOWER(name) LIKE ? OR LOWER(description) LIKE ?", fmt.Sprintf("%%%s%%", keyword), fmt.Sprintf("%%%s%%", keyword)).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductByID(ctx context.Context, productID int64) (*product.Product, error) {
	var productInfo product.Product
	// 这里假设 Product 表的主键是 ID 字段
	err := DB.WithContext(ctx).Where("id = ?", productID).First(&productInfo).Error
	if err != nil {
		return nil, err
	}
	return &productInfo, nil
}
