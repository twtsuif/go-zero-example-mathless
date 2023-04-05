package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mathless-backend/app/project/cmd/api/internal/config"
	"mathless-backend/app/project/cmd/rpc/projectservice"
)

type ServiceContext struct {
	Config     config.Config
	ProjectRpc projectservice.ProjectService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProjectRpc: projectservice.NewProjectService(zrpc.MustNewClient(c.ProjectRpcConf)),
	}
}
