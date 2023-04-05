package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProjectModel = (*customProjectModel)(nil)

type (
	// ProjectModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProjectModel.
	ProjectModel interface {
		projectModel
		FindProjectsByUid(ctx context.Context, uid int64) ([]*Project, error)
	}

	customProjectModel struct {
		*defaultProjectModel
	}
)

// NewProjectModel returns a model for the database table.
func NewProjectModel(conn sqlx.SqlConn, c cache.CacheConf) ProjectModel {
	return &customProjectModel{
		defaultProjectModel: newProjectModel(conn, c),
	}
}

func (m *defaultProjectModel) FindProjectsByUid(ctx context.Context, uid int64) ([]*Project, error) {
	selectBuilder := squirrel.Select(projectRows).From(m.table)
	sql, values, err := selectBuilder.Where("uid = ?", uid).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Project
	err = m.QueryRowsNoCacheCtx(ctx, &resp, sql, values...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
