package permission

import (
	"context"
	"storm-server/app/sys/cmd/rpc/client/permissionrpc"

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
	res, err := l.svcCtx.PermissionRpc.PermissionInfo(l.ctx, &permissionrpc.PermissionInfoReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.PermissionDetailResp{
		Id:         res.Id,
		Component:  res.Component,
		UpdateAt:   res.UpdateAt,
		Category:   res.Category,
		Code:       res.Code,
		Icon:       res.Icon,
		Status:     int(res.Status),
		Sort:       int(res.Sort),
		ParentId:   res.ParentId,
		UpdateUser: res.UpdateUser,
		Path:       res.Path,
		CreateUser: res.CreateUser,
		Name:       res.Name,
		CreateAt:   res.CreateAt,
	}, nil
}
