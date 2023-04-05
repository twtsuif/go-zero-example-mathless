package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mathless-backend/app/dependency/cmd/rpc/internal/config"
	"mathless-backend/app/dependency/model"
	"mathless-backend/common/tool/fc_sdk"
)

type ServiceContext struct {
	Config          config.Config
	FcTool          *fc_sdk.FcTool
	DependencyModel model.RequirementModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	aliyunConf := c.AliyunConf
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		FcTool: fc_sdk.NewFcTool(aliyunConf.FcConf.Endpoint,
			aliyunConf.FcConf.ApiVersion,
			aliyunConf.AccessKeyID,
			aliyunConf.AccessKeySecret),
		DependencyModel: model.NewRequirementModel(sqlConn, c.Cache),
	}
}
