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

type DepartmentAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDepartmentAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DepartmentAddLogic {
	return &DepartmentAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DepartmentAddLogic) DepartmentAdd(in *sysClient.AddDepartmentReq) (*sysClient.AddDepartmentResp, error) {
	var ancestors = "0"
	//父级id
	parentId := in.ParentId
	if parentId != "" && parentId != "0" {
		var parent sys.Department
		tx := l.svcCtx.Db.First(&parent, "id = ?", parentId)
		if tx.Error != nil {
			return nil, tx.Error
		}
		ancestors = parent.Ancestors + "," + parentId
	}

	department := sys.Department{
		ParentId:  parentId,
		Ancestors: ancestors,
		Name:      in.Name,
		Sort:      int(in.Sort),
	}
	result := l.svcCtx.Db.Create(&department)
	if result.Error != nil {
		err := errors.New("添加department失败:" + result.Error.Error())
		return nil, err
	}
	affected := result.RowsAffected
	return &sysClient.AddDepartmentResp{
		Data: strconv.FormatInt(affected, 10),
	}, nil
}
