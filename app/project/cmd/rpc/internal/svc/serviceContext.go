package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mathless-backend/app/project/cmd/rpc/internal/config"
	"mathless-backend/app/project/model"
)

type ServiceContext struct {
	Config       config.Config
	ProjectModel model.ProjectModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:       c,
		ProjectModel: model.NewProjectModel(conn, c.Cache),
	}
}
