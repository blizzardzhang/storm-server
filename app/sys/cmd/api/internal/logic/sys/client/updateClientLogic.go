package client

import (
	"context"
	"storm-server/app/sys/cmd/rpc/client/clientrpc"

	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateClientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateClientLogic {
	return &UpdateClientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateClientLogic) UpdateClient(req *types.UpdateClientReq) (resp *types.UpdateClientResp, err error) {
	res, err := l.svcCtx.ClientRpc.ClientUpdate(l.ctx, &clientrpc.UpdateClientReq{
		Id:                  req.Id,
		Name:                req.Name,
		GrantType:           req.GrantType,
		AccessTokenValidity: req.AccessTokenValidity,
		AdditionalInfo:      req.AdditionalInformation,
		Key:                 req.Key,
		RedirectUri:         req.RedirectUri,
		ClientId:            req.ClientId,
		Secret:              req.Secret,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateClientResp{
		Data: res.Data,
	}, nil
}
