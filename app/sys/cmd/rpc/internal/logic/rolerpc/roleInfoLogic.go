package rolerpclogic

import (
	"context"
	"storm-server/app/sys/model/sys"

	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleInfoLogic {
	return &RoleInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleInfoLogic) RoleInfo(in *sysClient.RoleInfoReq) (*sysClient.RoleInfoResp, error) {
	var role sys.Role
	result := l.svcCtx.Db.First(&role, "id = ?", in.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sysClient.RoleInfoResp{
		Id:       role.Id,
		Name:     role.Name,
		Code:     role.Code,
		Sort:     int64(role.Sort),
		Remark:   role.Remark,
		Type:     int64(role.Type),
		CreateAt: role.CreateAt.Format("2006-01-02 15:04:05"),
		UpdateAt: role.UpdateAt.Format("2006-01-02 15:04:05"),
	}, nil
}
