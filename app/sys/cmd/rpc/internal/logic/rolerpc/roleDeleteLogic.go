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

type RoleDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleDeleteLogic) RoleDelete(in *sysClient.DeleteRoleReq) (*sysClient.DeleteRoleResp, error) {
	result := l.svcCtx.Db.Delete(&sys.Role{}, in.Ids)
	if result.Error != nil {
		err := errors.New("删除角色失败: " + result.Error.Error())
		return nil, err
	}

	return &sysClient.DeleteRoleResp{
		Data: strconv.FormatInt(result.RowsAffected, 10),
	}, nil
}
