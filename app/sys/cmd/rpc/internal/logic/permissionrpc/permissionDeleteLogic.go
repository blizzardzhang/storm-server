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

type PermissionDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPermissionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionDeleteLogic {
	return &PermissionDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PermissionDeleteLogic) PermissionDelete(in *sysClient.DeletePermissionReq) (*sysClient.DeletePermissionResp, error) {
	ids := in.Ids
	//1.判断是否存在下级
	for _, id := range ids {
		//1.判断是否有下级
		var children []sys.Permission
		tx := l.svcCtx.Db.Where("parent_id = ?", id).Find(&children)
		if tx.Error != nil {
			err := errors.New("查询权限下级失败:" + tx.Error.Error())
			return nil, err
		}
		if len(children) > 0 {
			err := errors.New("存在下级，无法删除")
			return nil, err
		}

	}
	//2.执行删除
	result := l.svcCtx.Db.Delete(&sys.Permission{}, ids)
	if result.Error != nil {
		err := errors.New("删除permission失败: " + result.Error.Error())
		return nil, err
	}
	affected := result.RowsAffected
	return &sysClient.DeletePermissionResp{
		Data: strconv.FormatInt(affected, 10),
	}, nil
}
