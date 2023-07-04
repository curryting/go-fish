package svc

import (
	"fish/service/user/cmd/internal/config"
	"fish/service/user/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.FishUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewFishUserModel(conn),
	}
}
