// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package svc

import (
	"response"
	"user_api/internal/config"
	"user_api/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config     config.Config
	UsersModel model.UserApiModel
	DB         *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	db := response.InitGorm(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		UsersModel: model.NewUserApiModel(mysqlConn),
		DB:         db,
	}
}
