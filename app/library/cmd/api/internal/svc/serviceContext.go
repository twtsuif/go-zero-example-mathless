package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mathless-backend/app/library/cmd/api/internal/config"
	"mathless-backend/app/library/cmd/rpc/libraryservice"
)

type ServiceContext struct {
	Config     config.Config
	LibraryRpc libraryservice.LibraryService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		LibraryRpc: libraryservice.NewLibraryService(zrpc.MustNewClient(c.LibraryRpcConf)),
	}
}
