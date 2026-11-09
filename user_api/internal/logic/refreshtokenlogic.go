// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package logic

import (
	"context"
	"errors"
	"response"

	"user_api/internal/svc"
	"user_api/internal/types"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	token, err := jwt.ParseWithClaims(req.RefreshToken, jwt.MapClaims{},
		func(t *jwt.Token) (any, error) {
			return []byte(l.svcCtx.Config.RefreshSecret), nil
		})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired refresh token")
	}

	claims := token.Claims.(jwt.MapClaims)
	userIDClaim, ok := claims["userId"].(float64)
	if !ok {
		return nil, errors.New("invalid or expired refresh token")
	}
	username, ok := claims["username"].(string)
	if !ok {
		return nil, errors.New("invalid or expired refresh token")
	}
	userId := int64(userIDClaim)

	// user, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	// if err != nil {
	// 	return nil, err
	// }

	newAccess, err := response.GenerateAccessToken(
		l.svcCtx.Config.Auth.AccessSecret,
		l.svcCtx.Config.Auth.AccessExpire,
		userId,
		username,
		"role",
	)
	if err != nil {
		return nil, err
	}

	newRefresh, err := response.GenerateRefreshToken(
		l.svcCtx.Config.RefreshSecret,
		l.svcCtx.Config.RefreshExpire,
		userId,
		username,
	)
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		AccessToken:  newAccess,
		RefreshToken: newRefresh,
		ExpiresIn:    l.svcCtx.Config.Auth.AccessExpire,
	}, nil
}
