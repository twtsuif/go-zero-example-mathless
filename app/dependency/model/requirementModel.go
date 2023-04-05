package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RequirementModel = (*customRequirementModel)(nil)

type (
	// RequirementModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRequirementModel.
	RequirementModel interface {
		requirementModel
		FindRequirementPageByName(ctx context.Context, name string, page, pageSize int32) ([]*Requirement, error)
		FindRequirementCountByName(ctx context.Context, name string) (int32, error)
	}

	customRequirementModel struct {
		*defaultRequirementModel
	}
)

// NewRequirementModel returns a model for the database table.
func NewRequirementModel(conn sqlx.SqlConn, c cache.CacheConf) RequirementModel {
	return &customRequirementModel{
		defaultRequirementModel: newRequirementModel(conn, c),
	}
}

func (m *defaultRequirementModel) FindRequirementPageByName(ctx context.Context, name string, page, pageSize int32) ([]*Requirement, error) {
	// 构造页数
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	// 构造sql
	sql, values, err := squirrel.Select(requirementRows).From(m.table).Where("name LIKE ?", "%"+name+"%").OrderBy("LENGTH(name)").Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	// 查询
	var resp []*Requirement
	err = m.QueryRowsNoCacheCtx(ctx, &resp, sql, values...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultRequirementModel) FindRequirementCountByName(ctx context.Context, name string) (int32, error) {
	// 构造sql
	sql, values, err := squirrel.Select("COUNT(*)").From(m.table).Where("name like ?", "%"+name+"%").ToSql()
	if err != nil {
		return 0, err
	}

	// 查询
	var resp int32
	err = m.QueryRowNoCacheCtx(ctx, &resp, sql, values...)
	if err != nil {
		return 0, err
	}

	return resp, nil
}
