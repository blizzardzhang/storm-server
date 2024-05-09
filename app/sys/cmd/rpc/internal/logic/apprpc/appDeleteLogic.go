package apprpclogic

import (
	"context"
	"storm-server/app/sys/model/sys"
	"strconv"

	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppDeleteLogic {
	return &AppDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppDeleteLogic) AppDelete(in *sysClient.DeleteAppReq) (*sysClient.DeleteAppResp, error) {
	result := l.svcCtx.Db.Delete(&sys.App{}, in.Ids)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sysClient.DeleteAppResp{
		Data: strconv.FormatInt(result.RowsAffected, 10),
	}, nil
}
