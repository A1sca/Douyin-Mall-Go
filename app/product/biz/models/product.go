package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

// Product 产品模型
type Product struct {
	gorm.Model
	ID          uint32   `gorm:"column:id;primaryKey"`
	Name        string   `gorm:"column:name"`
	Description string   `gorm:"column:description"`
	Picture     string   `gorm:"column:picture"`
	Price       float32  `gorm:"column:price"`
	Categories  []string `gorm:"column:categories;type:text"`
}

func (Product) TableName() string {
	return "products"
}

func SearchProducts(ctx context.Context, db *gorm.DB, query string) ([]*Product, error) {
	var products []*Product
	// 构造模糊查询条件
	queryCondition := fmt.Sprintf("%%%s%%", query)
	// 使用 GORM 进行模糊查询
	if err := db.Where("name LIKE ? OR description LIKE ?", queryCondition, queryCondition).Find(&products).Error; err != nil {
		return nil, fmt.Errorf("failed to search products: %w", err)
	}
	return products, nil
}
