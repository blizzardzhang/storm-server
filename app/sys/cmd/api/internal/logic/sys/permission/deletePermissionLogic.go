package permission

import (
	"context"
	"storm-server/app/sys/cmd/rpc/client/permissionrpc"

	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePermissionLogic {
	return &DeletePermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePermissionLogic) DeletePermission(req *types.DeletePermissionReq) (resp *types.DeletePermissionResp, err error) {
	res, err := l.svcCtx.PermissionRpc.PermissionDelete(l.ctx, &permissionrpc.DeletePermissionReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}
	return &types.DeletePermissionResp{
		Data: res.Data,
	}, nil
}
