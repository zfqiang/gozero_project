// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package logic

import (
	"context"
	"errors"
	"response"
	"user_api/internal/svc"
	"user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	// userApiV2, err := l.svcCtx.DB.Table("user_api").Where("username = ?", req.Username).First()
	userApi, err := l.svcCtx.UsersModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		return nil, errors.New("账号不存在")
	}
	if userApi.Password != req.Password {
		return nil, errors.New("密码不正确")
	}

	accessToken, err := response.GenerateAccessToken(
		l.svcCtx.Config.Auth.AccessSecret,
		l.svcCtx.Config.Auth.AccessExpire,
		1,
		userApi.Username,
		"role",
	)
	if err != nil {
		return nil, err
	}

	refreshToken, err := response.GenerateRefreshToken(
		l.svcCtx.Config.RefreshSecret,
		l.svcCtx.Config.RefreshExpire,
		1,
		req.Username,
	)
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    l.svcCtx.Config.Auth.AccessExpire,
	}, nil
}
