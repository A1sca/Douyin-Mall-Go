package dal

import (
	"github.com/A1sca/Douyin-Mall-Go/app/payment/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
