package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mathless-backend/app/user/cmd/api/internal/config"
	"mathless-backend/app/user/cmd/rpc/userservice"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userservice.NewUserService(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
