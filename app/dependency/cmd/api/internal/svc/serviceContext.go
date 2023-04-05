package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mathless-backend/app/dependency/cmd/api/internal/config"
	"mathless-backend/app/dependency/cmd/rpc/dependencyservice"
	"mathless-backend/common/tool/fc_sdk"
	"mathless-backend/common/tool/oss_sdk"
)

type ServiceContext struct {
	Config        config.Config
	DependencyRpc dependencyservice.DependencyService
	FcTool        *fc_sdk.FcTool
	OssTool       *oss_sdk.OssTool
}

func NewServiceContext(c config.Config) *ServiceContext {
	aliyunConf := c.AliyunConf
	return &ServiceContext{
		Config:        c,
		DependencyRpc: dependencyservice.NewDependencyService(zrpc.MustNewClient(c.DependencyRpcConf)),
		FcTool: fc_sdk.NewFcTool(aliyunConf.FcConf.Endpoint,
			aliyunConf.FcConf.ApiVersion,
			aliyunConf.AccessKeyID,
			aliyunConf.AccessKeySecret),
		OssTool: oss_sdk.NewOssTool(aliyunConf.OssConf.Endpoint,
			aliyunConf.AccessKeyID,
			aliyunConf.AccessKeySecret,
			aliyunConf.OssConf.Bucket),
	}
}
