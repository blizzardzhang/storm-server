package permission

import (
	"context"
	"storm-server/app/sys/cmd/rpc/client/permissionrpc"

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
	res, err := l.svcCtx.PermissionRpc.PermissionAdd(l.ctx, &permissionrpc.AddPermissionReq{
		Name:      req.Name,
		Code:      req.Code,
		Component: req.Component,
		Icon:      req.Icon,
		Path:      req.Path,
		Sort:      int64(req.Sort),
		Category:  req.Category,
		ParentId:  req.ParentId,
	})
	if err != nil {
		return nil, err
	}

	return &types.SavePermissionResp{
		Data: res.Data,
	}, nil
}
