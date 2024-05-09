package permissionrpclogic

import (
	"context"
	"storm-server/app/sys/model/sys"

	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPermissionInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionInfoLogic {
	return &PermissionInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PermissionInfoLogic) PermissionInfo(in *sysClient.PermissionInfoReq) (*sysClient.PermissionInfoResp, error) {
	var permission sys.Permission
	if err := l.svcCtx.Db.Where("id = ?", in.Id).First(&permission).Error; err != nil {
		return nil, err
	}

	return &sysClient.PermissionInfoResp{
		Id:         permission.Id,
		ParentId:   permission.ParentId,
		Name:       permission.Name,
		Code:       permission.Code,
		Component:  permission.Component,
		Icon:       permission.Icon,
		Path:       permission.Path,
		Sort:       int64(permission.Sort),
		Status:     int64(permission.Status),
		Category:   permission.Category,
		CreateUser: permission.CreateBy,
		UpdateUser: permission.UpdateBy,
		CreateAt:   permission.CreateAt.Format("2006-01-02 15:04:05"),
		UpdateAt:   permission.UpdateAt.Format("2006-01-02 15:04:05"),
	}, nil
}
