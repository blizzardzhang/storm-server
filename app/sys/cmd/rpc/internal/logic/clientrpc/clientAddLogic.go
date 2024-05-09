package clientrpclogic

import (
	"context"
	"errors"
	"storm-server/app/sys/model/sys"
	"strconv"

	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClientAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClientAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClientAddLogic {
	return &ClientAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClientAddLogic) ClientAdd(in *sysClient.AddClientReq) (*sysClient.AddClientResp, error) {
	client := sys.Client{
		Name:                  in.Name,
		ClientId:              in.ClientId,
		Key:                   in.Key,
		Secret:                in.Secret,
		GrantType:             in.GrantType,
		RedirectUri:           in.RedirectUri,
		AdditionalInformation: in.AdditionalInfo,
		AccessTokenValidity:   in.AccessTokenValidity,
		RefreshTokenValidity:  in.RefreshTokenValidity,
	}
	result := l.svcCtx.Db.Create(&client) //指针数据
	if result.Error != nil {
		err := errors.New("添加app失败:" + result.Error.Error())
		return nil, err
	}
	affected := result.RowsAffected

	return &sysClient.AddClientResp{
		Data: strconv.FormatInt(affected, 10),
	}, nil
}
