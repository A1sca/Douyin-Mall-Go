package model

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type Address struct {
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	ZipCode       int32  `json:"zip_code"`
}

type OrderItem struct {
	ProductId  uint32  `json:"product_id"`
	ProductNum int32   `json:"product_num"`
	Cost       float32 `json:"cost"`
}

type Order struct {
	OrderId      string        `gorm:"size:255;not null;primary key"`
	UserId       uint32        `gorm:"not null"`
	UserCurrency string        `gorm:"size:255"`
	Email        string        `gorm:"size:255"`
	Address      Address       `gorm:"type:json"`
	OrderItems   OrderItemList `gorm:"type:json"`
}

// 指定用户表名为order, 不指定的话, 默认是orders
func (Order) TableName() string {
	return "order"
}

func Create(ctx context.Context, db *gorm.DB, order *Order) error {
	return db.WithContext(ctx).Create(order).Error
}

func GetOrderListByUserId(ctx context.Context, db *gorm.DB, userId uint32) (orders []Order, err error) {
	// err = db.Model(&Order{}).Where(&Order{UserId: userId}).Preload("OrderItems").Find(&orders).Error
	err = db.Model(&Order{}).Where("user_id = ?", userId).Find(&orders).Error
	return
}

// OrderItemList is a wrapper around []OrderItem for custom serialization.
type OrderItemList []OrderItem

// Value implements the driver.Valuer interface to serialize the OrderItemList to JSON.
func (o OrderItemList) Value() (driver.Value, error) {
	if len(o) == 0 {
		return nil, nil
	}
	return json.Marshal(o)
}

// Scan implements the sql.Scanner interface to deserialize the OrderItemList from JSON.
func (o *OrderItemList) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to convert value to []byte")
	}

	return json.Unmarshal(bytes, o)
}

func (a Address) Value() (driver.Value, error) {
	if a == (Address{}) { // 判断是否为空结构体
		return nil, nil
	}
	return json.Marshal(a)
}

// Scan implements the sql.Scanner interface to deserialize the Address from JSON.
func (a *Address) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to convert value to []byte")
	}

	return json.Unmarshal(bytes, (*Address)(a))
}
