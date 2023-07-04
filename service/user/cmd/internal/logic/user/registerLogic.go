package user

import (
	"context"
	"errors"
	"fish/service/user/cmd/internal/svc"
	"fish/service/user/cmd/internal/types"
	"fish/service/user/model"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.CommonRes, err error) {
	// todo: add your logic here and delete this line
	// 验证参数
	if len(strings.TrimSpace(req.Username)) == 0 {
		return nil, errors.New("参数错误")
	}
	_, err1 := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	switch err1 {
	case nil:
	case model.ErrNotFound:
		return nil, errors.New("用户名不存在")
	default:
		return nil, err
	}

	//if userInfo.Username != req.Username {
	//	return nil, errors.New("用户密码不正确")
	//}

	// ---start---
	//now := time.Now().Unix()
	//accessExpire := l.svcCtx.Config.Auth.AccessExpire
	//jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id)
	//if err != nil {
	//	return nil, err
	//}
	// ---end---

	return &types.CommonRes{
		IRet: 0,
		SMsg: req.Username + " register success",
	}, nil
}
