package apprpclogic

import (
	"context"
	"storm-server/app/sys/model/sys"

	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppInfoLogic {
	return &AppInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppInfoLogic) AppInfo(in *sysClient.AppInfoReq) (*sysClient.AppInfoResp, error) {
	var app sys.Client
	result := l.svcCtx.Db.First(&app, "id = ?", in.Id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &sysClient.AppInfoResp{
		Id:                   app.Id,
		AppId:                app.ClientId,
		Name:                 app.Name,
		Secret:               app.Secret,
		GrantType:            app.GrantType,
		RedirectUri:          app.RedirectUri,
		AdditionalInfo:       app.AdditionalInformation,
		AccessTokenValidity:  app.AccessTokenValidity,
		RefreshTokenValidity: app.RefreshTokenValidity,
		CreateAt:             app.CreateAt.Format("2006-01-02 15:04:05"),
		UpdateAt:             app.UpdateAt.Format("2006-01-02 15:04:05"),
		Status:               int64(app.Status),
	}, nil
}
