package clientrpclogic

import (
	"context"
	"storm-server/app/sys/model/sys"

	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClientInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientInfoLogic {
	return &ClientInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClientInfoLogic) ClientInfo(in *sysClient.ClientInfoReq) (*sysClient.ClientInfoResp, error) {
	var client sys.Client
	result := l.svcCtx.Db.First(&client, "id = ?", in.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &sysClient.ClientInfoResp{
		Id:                   client.Id,
		ClientId:             client.ClientId,
		Name:                 client.Name,
		Secret:               client.Secret,
		GrantType:            client.GrantType,
		RedirectUri:          client.RedirectUri,
		AdditionalInfo:       client.AdditionalInformation,
		AccessTokenValidity:  client.AccessTokenValidity,
		RefreshTokenValidity: client.RefreshTokenValidity,
		CreateAt:             client.CreateAt.Format("2006-01-02 15:04:05"),
		UpdateAt:             client.UpdateAt.Format("2006-01-02 15:04:05"),
		Status:               int64(client.Status),
	}, nil
}
