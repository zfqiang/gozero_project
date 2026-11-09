// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	RefreshSecret string
	RefreshExpire int64
	Mysql         struct {
		DataSource string
	}
	CacheRedis cache.CacheConf // redis缓存
}
