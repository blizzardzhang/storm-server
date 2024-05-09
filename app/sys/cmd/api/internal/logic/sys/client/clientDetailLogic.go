package client

import (
	"context"
	"storm-server/app/sys/cmd/rpc/client/clientrpc"

	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClientDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientDetailLogic {
	return &ClientDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClientDetailLogic) ClientDetail(req *types.ClientInfoReq) (resp *types.ClientInfoResp, err error) {
	info, err := l.svcCtx.ClientRpc.ClientInfo(l.ctx, &clientrpc.ClientInfoReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.ClientInfoResp{
		Id:                    info.Id,
		ClientId:              info.ClientId,
		Name:                  info.Name,
		Key:                   info.Key,
		RedirectUri:           info.RedirectUri,
		Secret:                info.Secret,
		GrantType:             info.GrantType,
		AccessTokenValidity:   info.AccessTokenValidity,
		AdditionalInformation: info.AdditionalInfo,
		Status:                int(info.Status),
		CreateUser:            info.CreateUser,
		CreateAt:              info.CreateAt,
		UpdateUser:            info.UpdateUser,
	}, nil
}
