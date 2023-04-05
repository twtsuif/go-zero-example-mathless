package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	DependencyRpcConf zrpc.RpcClientConf
	AliyunConf        struct {
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
}
