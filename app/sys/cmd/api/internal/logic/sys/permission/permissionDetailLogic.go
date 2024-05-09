package permission

import (
	"context"

	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPermissionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionDetailLogic {
	return &PermissionDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PermissionDetailLogic) PermissionDetail(req *types.PermissionDetailReq) (resp *types.PermissionDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
