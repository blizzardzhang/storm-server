package department

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
