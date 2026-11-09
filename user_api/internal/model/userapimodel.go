package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserApiModel = (*customUserApiModel)(nil)

type (
	// UserApiModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserApiModel.
	UserApiModel interface {
		userApiModel
		withSession(session sqlx.Session) UserApiModel
	}

	customUserApiModel struct {
		*defaultUserApiModel
	}
)

// NewUserApiModel returns a model for the database table.
func NewUserApiModel(conn sqlx.SqlConn) UserApiModel {
	return &customUserApiModel{
		defaultUserApiModel: newUserApiModel(conn),
	}
}

func (m *customUserApiModel) withSession(session sqlx.Session) UserApiModel {
	return NewUserApiModel(sqlx.NewSqlConnFromSession(session))
}
