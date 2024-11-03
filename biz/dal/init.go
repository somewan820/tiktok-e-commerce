package dal

import (
	"tiktok/biz/dal/mysql"
	"tiktok/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
