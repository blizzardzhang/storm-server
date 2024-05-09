package department

import (
	"context"
	"storm-server/app/sys/cmd/rpc/sysClient"

	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddDepartmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddDepartmentLogic {
	return &AddDepartmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddDepartmentLogic) AddDepartment(req *types.AddDepartmentReq) (resp *types.AddDepartmentResp, err error) {
	res, err := l.svcCtx.DepartmentRpc.DepartmentAdd(l.ctx, &sysClient.AddDepartmentReq{
		ParentId:  req.ParentId,
		Ancestors: req.Ancestors,
		Name:      req.Name,
		Sort:      int64(req.Sort),
	})
	if err != nil {
		return nil, err
	}
	return &types.AddDepartmentResp{
		Data: res.Data,
	}, nil
}
