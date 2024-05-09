package permissionrpclogic

import (
	"context"
	"errors"
	"storm-server/app/sys/model/sys"
	"strconv"

	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPermissionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionAddLogic {
	return &PermissionAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PermissionAddLogic) PermissionAdd(in *sysClient.AddPermissionReq) (*sysClient.AddPermissionResp, error) {
	permission := sys.Permission{
		ParentId:  in.ParentId,
		Name:      in.Name,
		Code:      in.Code,
		Component: in.Component,
		Icon:      in.Icon,
		Path:      in.Path,
		Sort:      int(in.Sort),
		Category:  in.Category,
	}
	result := l.svcCtx.Db.Create(&permission)
	if result.Error != nil {
		err := errors.New("添加department失败:" + result.Error.Error())
		return nil, err
	}
	affected := result.RowsAffected

	return &sysClient.AddPermissionResp{
		Data: strconv.FormatInt(affected, 10),
	}, nil
}
