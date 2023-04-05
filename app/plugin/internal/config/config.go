package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	AliyunConf struct {
		AccessKeyID     string
		AccessKeySecret string
		FcConf          struct {
			Endpoint   string
			ApiVersion string
		}
	}
}
