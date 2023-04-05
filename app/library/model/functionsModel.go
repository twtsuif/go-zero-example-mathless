package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FunctionsModel = (*customFunctionsModel)(nil)

type (
	// FunctionsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFunctionsModel.
	FunctionsModel interface {
		functionsModel
		FindFunctionsByLibraryId(ctx context.Context, LibraryId int64) ([]*Functions, error)
	}

	customFunctionsModel struct {
		*defaultFunctionsModel
	}
)

// NewFunctionsModel returns a model for the database table.
func NewFunctionsModel(conn sqlx.SqlConn, c cache.CacheConf) FunctionsModel {
	return &customFunctionsModel{
		defaultFunctionsModel: newFunctionsModel(conn, c),
	}
}

func (m *defaultFunctionsModel) FindFunctionsByLibraryId(ctx context.Context, LibraryId int64) ([]*Functions, error) {
	selectBuilder := squirrel.Select(functionsRows).From(m.table)
	sql, values, err := selectBuilder.Where("Library_id = ?", LibraryId).ToSql()
	if err != nil {
		return nil, err
	}
	var resp []*Functions
	err = m.QueryRowsNoCacheCtx(ctx, &resp, sql, values...)
	if err != nil {
		return nil, err
	}
	return resp, err
}
