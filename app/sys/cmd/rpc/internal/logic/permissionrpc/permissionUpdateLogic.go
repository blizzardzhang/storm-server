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

type PermissionUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPermissionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionUpdateLogic {
	return &PermissionUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PermissionUpdateLogic) PermissionUpdate(in *sysClient.UpdatePermissionReq) (*sysClient.UpdatePermissionResp, error) {
	updates := map[string]interface{}{
		"parent_id": in.ParentId,
		"name":      in.Name,
		"code":      in.Code,
		"path":      in.Path,
		"component": in.Component,
		"icon":      in.Icon,
		"sort":      in.Sort,
		"category":  in.Category,
	}
	var permission sys.Permission
	result := l.svcCtx.Db.Model(&permission).Where("id = ?", in.Id).Updates(updates)
	if result.Error != nil {
		err := errors.New("更新permission失败:" + result.Error.Error())
		return nil, err
	}

	return &sysClient.UpdatePermissionResp{
		Data: strconv.FormatInt(result.RowsAffected, 10),
	}, nil
}
