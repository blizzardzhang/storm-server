package departmentrpclogic

import (
	"context"
	"storm-server/app/sys/model/sys"

	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DepartmentInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDepartmentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DepartmentInfoLogic {
	return &DepartmentInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DepartmentInfoLogic) DepartmentInfo(in *sysClient.DepartmentInfoReq) (*sysClient.DepartmentInfoResp, error) {
	var department sys.Department
	result := l.svcCtx.Db.First(&department, "id = ?", in.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &sysClient.DepartmentInfoResp{
		Id:        department.Id,
		ParentId:  department.ParentId,
		Ancestors: department.Ancestors,
		Name:      department.Name,
		Sort:      int64(department.Sort),
		Status:    int64(department.Status),
		CreateAt:  department.CreateAt.Format("2006-01-02 15:04:05"),
		UpdateAt:  department.UpdateAt.Format("2006-01-02 15:04:05"),
	}, nil
}
