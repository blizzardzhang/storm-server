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

type RoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleUpdateLogic) RoleUpdate(in *sysClient.UpdateRoleReq) (*sysClient.UpdateRoleResp, error) {
	updates := map[string]interface{}{
		"type":   in.Type,
		"name":   in.Name,
		"code":   in.Code,
		"sort":   in.Sort,
		"remark": in.Remark,
	}
	var role sys.Role
	result := l.svcCtx.Db.Model(&role).Where("id = ?", in.Id).Updates(updates)
	if result.Error != nil {
		err := errors.New("更新失败:" + result.Error.Error())
		return nil, err
	}

	return &sysClient.UpdateRoleResp{
		Data: strconv.FormatInt(result.RowsAffected, 10),
	}, nil
}
