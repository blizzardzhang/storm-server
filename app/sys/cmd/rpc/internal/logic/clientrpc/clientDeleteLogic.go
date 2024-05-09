package clientrpclogic

import (
	"context"
	"storm-server/app/sys/model/sys"
	"strconv"

	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClientDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientDeleteLogic {
	return &ClientDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClientDeleteLogic) ClientDelete(in *sysClient.DeleteClientReq) (*sysClient.DeleteClientResp, error) {
	result := l.svcCtx.Db.Delete(&sys.Client{}, in.Ids)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sysClient.DeleteClientResp{
		Data: strconv.FormatInt(result.RowsAffected, 10),
	}, nil
}
