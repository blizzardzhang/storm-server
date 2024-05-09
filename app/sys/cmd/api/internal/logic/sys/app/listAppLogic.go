package app

import (
	"context"

	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAppLogic {
	return &ListAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListAppLogic) ListApp(req *types.ListAppReq) (resp *types.ListAppResp, err error) {
	// todo: add your logic here and delete this line

	return
}
