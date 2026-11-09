// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package logic

import (
	"context"
	"encoding/json"
	"errors"

	"user_api/internal/svc"
	"user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResp, err error) {
	userIDClaim, ok := l.ctx.Value("userId").(json.Number)
	if !ok {
		return nil, errors.New("user ID claim not found")
	}
	userId, err := userIDClaim.Int64()
	if err != nil {
		return nil, errors.New("invalid user ID claim")
	}

	role, ok := l.ctx.Value("role").(string)
	if !ok {
		return nil, errors.New("role claim not found")
	}

	userApi, err := l.svcCtx.UsersModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResp{
		ID:   userId,
		Name: userApi.Username,
		Role: role,
	}, nil
}
