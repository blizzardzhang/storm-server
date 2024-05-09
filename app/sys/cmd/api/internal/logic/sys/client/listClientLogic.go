package client

import (
	"context"
	"github.com/jinzhu/copier"
	"storm-server/app/sys/cmd/rpc/client/clientrpc"

	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListClientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListClientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListClientLogic {
	return &ListClientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListClientLogic) ListClient(req *types.ListClientReq) (resp *types.ListClientResp, err error) {
	res, err := l.svcCtx.ClientRpc.ClientList(l.ctx, &clientrpc.ListClientReq{
		Current:  int64(req.Current),
		PageSize: int64(req.PageSize),
		Name:     req.Name,
	})
	if err != nil {
		return nil, err
	}
	var list []types.ClientInfoResp
	for _, item := range res.List {
		var app types.ClientInfoResp
		_ = copier.Copy(&app, item)
		list = append(list, app)
	}
	return &types.ListClientResp{
		Records:   list,
		Current:   int(res.Current),
		PageSize:  int(res.PageSize),
		Total:     int(res.Total),
		TotalPage: int(res.TotalPage),
	}, nil
}
