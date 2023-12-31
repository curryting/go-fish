// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	fishUserFieldNames          = builder.RawFieldNames(&FishUser{})
	fishUserRows                = strings.Join(fishUserFieldNames, ",")
	fishUserRowsExpectAutoSet   = strings.Join(stringx.Remove(fishUserFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), ",")
	fishUserRowsWithPlaceHolder = strings.Join(stringx.Remove(fishUserFieldNames, "`id`", "`updated_at`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`"), "=?,") + "=?"
)

type (
	fishUserModel interface {
		Insert(ctx context.Context, data *FishUser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*FishUser, error)
		FindOneByUsername(ctx context.Context, username string) (*FishUser, error)
		Update(ctx context.Context, data *FishUser) error
		Delete(ctx context.Context, id int64) error
	}

	defaultFishUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	FishUser struct {
		Id         int64     `db:"id"`
		Username   string    `db:"username"` // username
		Password   string    `db:"password"` // password
		Gender     int64     `db:"gender"`   // 0-un status｜1-male｜2-female
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newFishUserModel(conn sqlx.SqlConn) *defaultFishUserModel {
	return &defaultFishUserModel{
		conn:  conn,
		table: "`fish_user`",
	}
}

func (m *defaultFishUserModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultFishUserModel) FindOne(ctx context.Context, id int64) (*FishUser, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", fishUserRows, m.table)
	var resp FishUser
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFishUserModel) FindOneByUsername(ctx context.Context, username string) (*FishUser, error) {
	var resp FishUser
	query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", fishUserRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFishUserModel) Insert(ctx context.Context, data *FishUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, fishUserRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Username, data.Password, data.Gender)
	return ret, err
}

func (m *defaultFishUserModel) Update(ctx context.Context, newData *FishUser) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, fishUserRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Username, newData.Password, newData.Gender, newData.Id)
	return err
}

func (m *defaultFishUserModel) tableName() string {
	return m.table
}
