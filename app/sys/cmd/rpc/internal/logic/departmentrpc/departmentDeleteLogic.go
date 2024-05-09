package departmentrpclogic

import (
	"context"
	"errors"
	"storm-server/app/sys/model/sys"
	"strconv"

	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DepartmentDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDepartmentDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DepartmentDeleteLogic {
	return &DepartmentDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DepartmentDeleteLogic) DepartmentDelete(in *sysClient.DeleteDepartmentReq) (*sysClient.DeleteDepartmentResp, error) {
	ids := in.Ids
	for _, id := range ids {
		// 1.判断是否有下级
		var children []sys.Department
		tx := l.svcCtx.Db.Where("parent_id = ?", id).Find(&children)
		if tx.Error != nil {
			err := errors.New("查询下级失败:" + tx.Error.Error())
			return nil, err
		}
		if len(children) > 0 {
			err := errors.New("存在下级,无法删除")
			return nil, err
		}
		//2.判断是否有用户
		var users []sys.User
		find := l.svcCtx.Db.Where("dept_id = ?", id).Find(&users)
		if find.Error != nil {
			err := errors.New("查询机构下用户失败:" + tx.Error.Error())
			return nil, err
		}
		if len(users) > 0 {
			err := errors.New("存在用户,无法删除")
			return nil, err
		}
	}
	//3.执行删除
	result := l.svcCtx.Db.Delete(&sys.Department{}, ids)
	if result.Error != nil {
		err := errors.New("删除department失败:" + result.Error.Error())
		return nil, err
	}
	affected := result.RowsAffected

	return &sysClient.DeleteDepartmentResp{
		Data: strconv.FormatInt(affected, 10),
	}, nil
}
