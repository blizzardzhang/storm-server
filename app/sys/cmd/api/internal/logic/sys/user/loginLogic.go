package user

import (
	"context"
	"github.com/jinzhu/copier"
	"storm-server/app/sys/cmd/rpc/client/userrpc"

	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"

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

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResp, err error) {
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &userrpc.LoginReq{
		AppId:    req.AppId,
		Account:  req.Account,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	var user types.UserInfo
	_ = copier.Copy(&user, res.UserInfo)

	return &types.LoginResp{
		AccessToken:  res.AccessToken,
		AccessExpire: res.AccessExpire,
		RefreshAfter: res.RefreshAfter,
		User:         user,
		Permissions:  res.Permissions,
	}, nil
}
