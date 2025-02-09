package dal

import (
	"github.com/A1sca/Douyin-Mall-Go/app/cart/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
