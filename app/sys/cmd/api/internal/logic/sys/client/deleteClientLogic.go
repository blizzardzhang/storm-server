package client

import (
	"context"
	"storm-server/app/sys/cmd/rpc/client/clientrpc"

	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteClientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteClientLogic {
	return &DeleteClientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteClientLogic) DeleteClient(req *types.DeleteClientReq) (resp *types.DeleteClientResp, err error) {
	res, err := l.svcCtx.ClientRpc.ClientDelete(l.ctx, &clientrpc.DeleteClientReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}
	return &types.DeleteClientResp{
		Data: res.Data,
	}, nil
}
