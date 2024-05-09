package department

import (
	"context"
	"storm-server/app/sys/cmd/rpc/client/departmentrpc"

	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DepartmentInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDepartmentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DepartmentInfoLogic {
	return &DepartmentInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DepartmentInfoLogic) DepartmentInfo(req *types.DepartmentInfoReq) (resp *types.DepartmentInfoResp, err error) {
	info, err := l.svcCtx.DepartmentRpc.DepartmentInfo(l.ctx, &departmentrpc.DepartmentInfoReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.DepartmentInfoResp{
		Id:         info.Id,
		Name:       info.Name,
		Sort:       int(info.Sort),
		CreateAt:   info.CreateAt,
		CreateUser: info.CreateUser,
		UpdateAt:   info.UpdateAt,
		UpdateUser: info.UpdateUser,
		Ancestors:  info.Ancestors,
	}, nil
}
