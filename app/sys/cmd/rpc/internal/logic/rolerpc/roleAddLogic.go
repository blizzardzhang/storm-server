package rolerpclogic

import (
	"context"
	"errors"
	"storm-server/app/sys/model/sys"
	"strconv"

	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddLogic {
	return &RoleAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleAddLogic) RoleAdd(in *sysClient.AddRoleReq) (*sysClient.AddRoleResp, error) {
	role := sys.Role{
		Type:   int(in.Type),
		Name:   in.Name,
		Code:   in.Code,
		Sort:   int(in.Sort),
		Remark: in.Remark,
	}
	result := l.svcCtx.Db.Create(&role) //指针数据
	if result.Error != nil {
		err := errors.New("添加失败:" + result.Error.Error())
		return nil, err
	}
	affected := result.RowsAffected
	return &sysClient.AddRoleResp{
		Data: strconv.FormatInt(affected, 10),
	}, nil
}
