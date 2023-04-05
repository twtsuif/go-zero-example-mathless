package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	AliyunConf struct {
		AccessKeyID     string
		AccessKeySecret string
		FcConf          struct {
			Endpoint   string
			ApiVersion string
		}
		OssConf struct {
			Endpoint string
			Bucket   string
		}
	}
	DB struct {
		DataSource string
	}
	Cache cache.CacheConf
}
