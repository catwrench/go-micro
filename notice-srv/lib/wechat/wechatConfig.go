package wechat

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"notice-srv/lib"
)

//初始化微信配置
func InitWechat() *wechat.Wechat {
	wc := wechat.NewWechat()
	redisOpts := &cache.RedisOpts{
		Host:        lib.Config.Redis.Host,
		Password:    lib.Config.Redis.Password,
		Database:    lib.Config.Redis.Database,
		MaxActive:   lib.Config.Redis.MaxActive,
		MaxIdle:     lib.Config.Redis.MaxIdle,
		IdleTimeout: lib.Config.Redis.IdleTimeout,
	}
	redisCache := cache.NewRedis(redisOpts)
	wc.SetCache(redisCache)
	return wc
}
