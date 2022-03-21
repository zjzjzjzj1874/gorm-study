package helper

import (
	"github.com/go-redis/redis/v8"
	"github.com/zjzjzjzj1874/gorm-study/cmd/global"
)

var Redis *redis.Client

func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     global.GlobalConfig.RedisInfo.Addr,
		Password: global.GlobalConfig.RedisInfo.Pwd,
		DB:       global.GlobalConfig.RedisInfo.DB,
	})
}
