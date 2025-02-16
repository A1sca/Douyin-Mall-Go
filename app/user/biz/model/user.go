package model

import (
	"context"

	"gorm.io/gorm"
)

type Address struct {
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	Detail   string `json:"detail"`
}

type User struct {
	gorm.Model
	Username       string    `gorm:"size:255;not null"`
	Phone          string    `gorm:"size:255"`
	Gender         int32     `gorm:"not null"`
	Age            int32     `gorm:"not null"`
	Avatar         string    `gorm:"size:255"`
	Email          string    `gorm:"size:255"`
	Balance        int32     `gorm:"not null"`
	Desc           string    `gorm:"size:1024"`
	Addresses      []Address `gorm:"type:json"`
	PasswordHashed string    `gorm:"size:255;not null"`
}

// 指定用户表名为user, 不指定的话, 默认是users
func (User) TableName() string {
	return "user"
}

func Create(ctx context.Context, db *gorm.DB, user *User) error {
	return db.WithContext(ctx).Create(user).Error
}

func GetByEmail(ctx context.Context, db *gorm.DB, email string) (*User, error) {
	var user User
	err := db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}

func GetById(ctx context.Context, db *gorm.DB, id string) (*User, error) {
	var user User
	err := db.WithContext(ctx).Where("user_id = ?", id).Find(&user).Error
	return &user, err
}

func DeleteById(ctx context.Context, db *gorm.DB, id string) (error) {
	err := db.WithContext(ctx).Delete(User{}, "user_id = ?", id)
	return err.Error
}