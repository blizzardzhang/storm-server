package role

import (
	"context"
	"storm-server/app/sys/cmd/rpc/client/rolerpc"

	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleLogic {
	return &GetRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleLogic) GetRole(req *types.RoleInfoReq) (resp *types.RoleInfoResp, err error) {
	info, err := l.svcCtx.RoleRpc.RoleInfo(l.ctx, &rolerpc.RoleInfoReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.RoleInfoResp{
		Id:         info.Id,
		Name:       info.Name,
		Code:       info.Code,
		Sort:       int(info.Sort),
		Remark:     info.Remark,
		Type:       int(info.Type),
		Status:     int(info.Status),
		CreateAt:   info.CreateAt,
		CreateUser: info.CreateUser,
		UpdateAt:   info.UpdateAt,
		UpdateUser: info.UpdateUser,
	}, nil
}
