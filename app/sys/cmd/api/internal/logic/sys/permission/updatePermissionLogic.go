package permission

import (
	"context"
	"storm-server/app/sys/cmd/rpc/client/permissionrpc"

	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePermissionLogic {
	return &UpdatePermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePermissionLogic) UpdatePermission(req *types.UpdatePermissionReq) (resp *types.UpdatePermissionResp, err error) {
	res, err := l.svcCtx.PermissionRpc.PermissionUpdate(l.ctx, &permissionrpc.UpdatePermissionReq{
		Id:        req.Id,
		Sort:      int64(req.Sort),
		Code:      req.Code,
		Category:  req.Category,
		ParentId:  req.ParentId,
		Component: req.Component,
		Path:      req.Path,
		Icon:      req.Icon,
		Name:      req.Name,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdatePermissionResp{
		Data: res.Data,
	}, nil
}
