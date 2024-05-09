package departmentrpclogic

import (
	"context"
	"storm-server/app/sys/model/sys"

	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DepartmentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDepartmentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DepartmentListLogic {
	return &DepartmentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DepartmentListLogic) DepartmentList(in *sysClient.ListDepartmentReq) (*sysClient.ListDepartmentResp, error) {
	//1.查询所有
	var departments []sys.Department
	tx := l.svcCtx.Db.Find(&departments)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var result []*sysClient.DepartmentInfoResp
	for _, department := range departments {
		result = append(result, &sysClient.DepartmentInfoResp{
			Id:         department.Id,
			ParentId:   department.ParentId,
			Ancestors:  department.Ancestors,
			Name:       department.Name,
			Sort:       int64(department.Sort),
			CreateAt:   department.CreateAt.Format("2006-01-02 15:04:05"),
			UpdateAt:   department.UpdateAt.Format("2006-01-02 15:04:05"),
			CreateUser: department.CreateBy,
			UpdateUser: department.UpdateBy,
		})
	}

	return &sysClient.ListDepartmentResp{
		List: result,
	}, nil
}
