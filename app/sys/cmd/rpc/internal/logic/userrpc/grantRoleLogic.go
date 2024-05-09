package userrpclogic

import (
	"context"

	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrantRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGrantRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrantRoleLogic {
	return &GrantRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GrantRoleLogic) GrantRole(in *sysClient.GrantRoleReq) (*sysClient.GrantRoleResp, error) {
	// todo: add your logic here and delete this line

	return &sysClient.GrantRoleResp{}, nil
}
