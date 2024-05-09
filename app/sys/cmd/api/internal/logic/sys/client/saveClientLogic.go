package client

import (
	"context"
	"storm-server/app/sys/cmd/rpc/client/clientrpc"

	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveClientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveClientLogic {
	return &SaveClientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveClientLogic) SaveClient(req *types.AddClientReq) (resp *types.AddClientResp, err error) {
	//userId := l.ctx.Value("userId").(string)
	res, err := l.svcCtx.ClientRpc.ClientAdd(l.ctx, &clientrpc.AddClientReq{
		ClientId:             req.ClientId,
		RedirectUri:          req.RedirectUri,
		AccessTokenValidity:  req.AccessTokenValidity,
		Key:                  req.Key,
		Secret:               req.Secret,
		RefreshTokenValidity: req.RefreshTokenValidity,
		Name:                 req.Name,
		GrantType:            req.GrantType,
		AdditionalInfo:       req.AdditionalInformation,
	})
	if err != nil {
		return nil, err
	}

	return &types.AddClientResp{
		Data: res.Data,
	}, nil
}
