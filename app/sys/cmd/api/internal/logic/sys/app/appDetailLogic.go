package app

import (
	"context"

	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppDetailLogic {
	return &AppDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppDetailLogic) AppDetail(req *types.AppInfoReq) (resp *types.AppInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
