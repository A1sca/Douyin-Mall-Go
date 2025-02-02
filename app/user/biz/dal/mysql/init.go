package mysql

import (
	"github.com/A1sca/Douyin-Mall-Go/app/user/biz/model"
	"github.com/A1sca/Douyin-Mall-Go/app/user/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		// DB, err = gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/gomall?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	// if conf.GetConf().Env != "online" {
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	// }
}
