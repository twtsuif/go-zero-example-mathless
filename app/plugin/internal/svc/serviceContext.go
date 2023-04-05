package svc

import (
	"mathless-backend/app/plugin/internal/config"
	"mathless-backend/common/tool/fc_sdk"
	"mathless-backend/common/tool/oss_sdk"
)

type ServiceContext struct {
	Config  config.Config
	FcTool  *fc_sdk.FcTool
	OssTool *oss_sdk.OssTool
}

func NewServiceContext(c config.Config) *ServiceContext {
	aliyunConf := c.AliyunConf
	return &ServiceContext{
		Config: c,
		FcTool: fc_sdk.NewFcTool(aliyunConf.FcConf.Endpoint,
			aliyunConf.FcConf.ApiVersion,
			aliyunConf.AccessKeyID,
			aliyunConf.AccessKeySecret),
	}
}
