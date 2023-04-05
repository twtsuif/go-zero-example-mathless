package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LibraryModel = (*customLibraryModel)(nil)

type (
	// LibraryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLibraryModel.
	LibraryModel interface {
		libraryModel
		FindLibrariesByUid(ctx context.Context, uid int64) ([]*Library, error)
		ChangeLibraryStatus(ctx context.Context, libraryId int64, status int64) error
	}

	customLibraryModel struct {
		*defaultLibraryModel
	}
)

// NewLibraryModel returns a model for the database table.
func NewLibraryModel(conn sqlx.SqlConn, c cache.CacheConf) LibraryModel {
	return &customLibraryModel{
		defaultLibraryModel: newLibraryModel(conn, c),
	}
}

func (m *defaultLibraryModel) FindLibrariesByUid(ctx context.Context, uid int64) ([]*Library, error) {
	selectBuilder := squirrel.Select(libraryRows).From(m.table)
	sql, values, err := selectBuilder.Where("uid = ? ", uid).ToSql()
	if err != nil {
		return nil, err
	}
	var resp []*Library
	err = m.QueryRowsNoCacheCtx(ctx, &resp, sql, values...)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (m *defaultLibraryModel) ChangeLibraryStatus(ctx context.Context, libraryId int64, status int64) error {
	libraryModel, err := m.FindOne(ctx, libraryId)
	if err != nil {
		return err
	}

	libraryModel.Status = status
	return m.Update(ctx, libraryModel)
}
