package permissionrpclogic

import (
	"context"
	"storm-server/app/sys/model/sys"

	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPermissionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionListLogic {
	return &PermissionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PermissionListLogic) PermissionList(in *sysClient.ListPermissionReq) (*sysClient.ListPermissionResp, error) {
	var permissions []sys.Permission
	if err := l.svcCtx.Db.Find(&permissions).Error; err != nil {
		return nil, err
	}

	var result []*sysClient.PermissionInfoResp
	for _, permission := range permissions {
		result = append(result, &sysClient.PermissionInfoResp{
			Id:         permission.Id,
			Icon:       permission.Icon,
			Path:       permission.Path,
			Sort:       int64(permission.Sort),
			Category:   permission.Category,
			ParentId:   permission.ParentId,
			Name:       permission.Name,
			Code:       permission.Code,
			Component:  permission.Component,
			Status:     int64(permission.Status),
			CreateUser: permission.CreateBy,
			UpdateUser: permission.UpdateBy,
			CreateAt:   permission.CreateAt.Format("2006-01-02 15:04:05"),
			UpdateAt:   permission.UpdateAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &sysClient.ListPermissionResp{
		List: result,
	}, nil
}
