package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	SMTPConf struct {
		Host     string
		Port     int
		Username string
		Password string
	}
	DB struct {
		DataSource string
	}
	Cache cache.CacheConf
}
