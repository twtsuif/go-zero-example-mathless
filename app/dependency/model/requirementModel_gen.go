// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	requirementFieldNames          = builder.RawFieldNames(&Requirement{})
	requirementRows                = strings.Join(requirementFieldNames, ",")
	requirementRowsExpectAutoSet   = strings.Join(stringx.Remove(requirementFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	requirementRowsWithPlaceHolder = strings.Join(stringx.Remove(requirementFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheMathlessRequirementIdPrefix   = "cache:mathless:requirement:id:"
	cacheMathlessRequirementNamePrefix = "cache:mathless:requirement:name:"
)

type (
	requirementModel interface {
		Insert(ctx context.Context, data *Requirement) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Requirement, error)
		FindOneByName(ctx context.Context, name string) (*Requirement, error)
		Update(ctx context.Context, data *Requirement) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRequirementModel struct {
		sqlc.CachedConn
		table string
	}

	Requirement struct {
		Id   int64  `db:"id"`
		Name string `db:"name"`
	}
)

func newRequirementModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultRequirementModel {
	return &defaultRequirementModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`requirement`",
	}
}

func (m *defaultRequirementModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	mathlessRequirementIdKey := fmt.Sprintf("%s%v", cacheMathlessRequirementIdPrefix, id)
	mathlessRequirementNameKey := fmt.Sprintf("%s%v", cacheMathlessRequirementNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, mathlessRequirementIdKey, mathlessRequirementNameKey)
	return err
}

func (m *defaultRequirementModel) FindOne(ctx context.Context, id int64) (*Requirement, error) {
	mathlessRequirementIdKey := fmt.Sprintf("%s%v", cacheMathlessRequirementIdPrefix, id)
	var resp Requirement
	err := m.QueryRowCtx(ctx, &resp, mathlessRequirementIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", requirementRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRequirementModel) FindOneByName(ctx context.Context, name string) (*Requirement, error) {
	mathlessRequirementNameKey := fmt.Sprintf("%s%v", cacheMathlessRequirementNamePrefix, name)
	var resp Requirement
	err := m.QueryRowIndexCtx(ctx, &resp, mathlessRequirementNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", requirementRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, name); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRequirementModel) Insert(ctx context.Context, data *Requirement) (sql.Result, error) {
	mathlessRequirementIdKey := fmt.Sprintf("%s%v", cacheMathlessRequirementIdPrefix, data.Id)
	mathlessRequirementNameKey := fmt.Sprintf("%s%v", cacheMathlessRequirementNamePrefix, data.Name)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?)", m.table, requirementRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name)
	}, mathlessRequirementIdKey, mathlessRequirementNameKey)
	return ret, err
}

func (m *defaultRequirementModel) Update(ctx context.Context, newData *Requirement) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	mathlessRequirementIdKey := fmt.Sprintf("%s%v", cacheMathlessRequirementIdPrefix, data.Id)
	mathlessRequirementNameKey := fmt.Sprintf("%s%v", cacheMathlessRequirementNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, requirementRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Name, newData.Id)
	}, mathlessRequirementIdKey, mathlessRequirementNameKey)
	return err
}

func (m *defaultRequirementModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheMathlessRequirementIdPrefix, primary)
}

func (m *defaultRequirementModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", requirementRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultRequirementModel) tableName() string {
	return m.table
}
