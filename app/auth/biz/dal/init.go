package dal

import (
	"github.com/A1sca/Douyin-Mall-Go/app/auth/biz/dal/mysql"
	"github.com/A1sca/Douyin-Mall-Go/app/auth/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
