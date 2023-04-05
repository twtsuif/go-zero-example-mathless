package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mathless-backend/app/library/cmd/rpc/internal/config"
	"mathless-backend/app/library/model"
)

type ServiceContext struct {
	Config        config.Config
	LibraryModel  model.LibraryModel
	FunctionModel model.FunctionsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:        c,
		LibraryModel:  model.NewLibraryModel(sqlConn, c.Cache),
		FunctionModel: model.NewFunctionsModel(sqlConn, c.Cache),
	}
}
