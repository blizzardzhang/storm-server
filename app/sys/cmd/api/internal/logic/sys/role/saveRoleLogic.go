package role

import (
	"context"
	"storm-server/app/sys/cmd/rpc/client/rolerpc"

	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveRoleLogic {
	return &SaveRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveRoleLogic) SaveRole(req *types.RoleAddReq) (resp *types.RoleAddResp, err error) {
	res, err := l.svcCtx.RoleRpc.RoleAdd(l.ctx, &rolerpc.AddRoleReq{
		Name:   req.Name,
		Code:   req.Code,
		Sort:   int64(req.Sort),
		Remark: req.Remark,
		Type:   int64(req.Type),
	})
	if err != nil {
		return nil, err
	}
	return &types.RoleAddResp{
		Data: res.Data,
	}, nil
}
