package model

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func init() {
	var err error
	// 使用测试数据库连接
	dsn := "root:rootroot@tcp(127.0.0.1:3306)/gomall_test?charset=utf8mb4&parseTime=True&loc=Local"
	testDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	// 自动迁移创建测试表
	err = testDB.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
}

func TestUserCRUD(t *testing.T) {
	ctx := context.Background()
	
	// 测试创建用户
	testUser := &User{
		Username:       "testuser",
		Phone:         "13800138000",
		Gender:        1,
		Age:           20,
		Email:         "test@example.com",
		Balance:       100,
		PasswordHashed: "hashedpassword",
	}
	
	err := Create(ctx, testDB, testUser)
	assert.Nil(t, err)
	assert.NotZero(t, testUser.ID)
	
	// 将uint类型的ID转换为string
	userID := strconv.FormatUint(uint64(testUser.ID), 10)
	
	// 测试通过ID获取用户
	user, err := GetById(ctx, testDB, userID)
	assert.Nil(t, err)
	assert.Equal(t, testUser.Username, user.Username)
	
	// 测试通过邮箱获取用户
	user, err = GetByEmail(ctx, testDB, testUser.Email)
	assert.Nil(t, err)
	assert.Equal(t, testUser.Email, user.Email)
	
	// 测试更新用户
	updates := &User{
		Username: "updateduser",
		Age:     21,
	}
	err = UpdateById(ctx, testDB, userID, updates)
	assert.Nil(t, err)
	
	// 验证更新结果
	updatedUser, err := GetById(ctx, testDB, userID)
	assert.Nil(t, err)
	assert.Equal(t, "updateduser", updatedUser.Username)
	assert.Equal(t, int32(21), updatedUser.Age)
	
	// 测试删除用户
	err = DeleteById(ctx, testDB, userID)
	assert.Nil(t, err)
	
	// 验证删除结果
	deletedUser, err := GetById(ctx, testDB, userID)
	assert.Nil(t, err)
	assert.Nil(t, deletedUser)
} 