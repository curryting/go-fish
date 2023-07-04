package model

import (
	"github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FishUserModel = (*customFishUserModel)(nil)

type (
	// FishUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFishUserModel.
	FishUserModel interface {
		fishUserModel
	}

	customFishUserModel struct {
		*defaultFishUserModel
	}
)

// NewFishUserModel returns a model for the database table.
func NewFishUserModel(conn sqlx.SqlConn) FishUserModel {
	return &customFishUserModel{
		defaultFishUserModel: newFishUserModel(conn),
	}
}
