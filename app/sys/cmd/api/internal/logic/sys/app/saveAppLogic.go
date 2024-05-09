package app

import (
	"context"

	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveAppLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveAppLogic {
	return &SaveAppLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveAppLogic) SaveApp(req *types.AddAppReq) (resp *types.AddAppResp, err error) {
	// todo: add your logic here and delete this line

	return
}
