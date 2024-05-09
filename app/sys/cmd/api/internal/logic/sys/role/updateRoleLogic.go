package role

import (
	"context"
	"storm-server/app/sys/cmd/rpc/client/rolerpc"

	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.RoleUpdateReq) (resp *types.RoleUpdateResp, err error) {
	res, err := l.svcCtx.RoleRpc.RoleUpdate(l.ctx, &rolerpc.UpdateRoleReq{
		Id:     req.Id,
		Name:   req.Name,
		Code:   req.Code,
		Sort:   int64(req.Sort),
		Remark: req.Remark,
	})
	if err != nil {
		return nil, err
	}
	return &types.RoleUpdateResp{
		Data: res.Data,
	}, nil
}
