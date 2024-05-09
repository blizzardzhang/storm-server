package permission

import (
	"context"

	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SavePermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSavePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SavePermissionLogic {
	return &SavePermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SavePermissionLogic) SavePermission(req *types.SavePermissionReq) (resp *types.SavePermissionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
