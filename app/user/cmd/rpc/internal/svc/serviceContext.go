package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mathless-backend/app/user/cmd/rpc/internal/config"
	"mathless-backend/app/user/model"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	UserModel   model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:      c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {}),
		UserModel:   model.NewUserModel(sqlConn, c.Cache),
	}
}
