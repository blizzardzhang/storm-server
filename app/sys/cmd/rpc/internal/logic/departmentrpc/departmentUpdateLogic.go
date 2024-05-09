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

type DepartmentUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDepartmentUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DepartmentUpdateLogic {
	return &DepartmentUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DepartmentUpdateLogic) DepartmentUpdate(in *sysClient.UpdateDepartmentReq) (*sysClient.UpdateDepartmentResp, error) {
	updates := map[string]interface{}{
		"parent_id": in.ParentId,
		"ancestors": in.Ancestors,
		"name":      in.Name,
		"sort":      in.Sort,
	}
	var department sys.Department
	result := l.svcCtx.Db.Model(&department).Where("id = ?", in.Id).Updates(updates)
	if result.Error != nil {
		err := errors.New("更新失败:" + result.Error.Error())
		return nil, err
	}

	return &sysClient.UpdateDepartmentResp{
		Data: strconv.FormatInt(result.RowsAffected, 10),
	}, nil
}
