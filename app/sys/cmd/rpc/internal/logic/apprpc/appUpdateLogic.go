package apprpclogic

import (
	"context"
	"errors"
	"storm-server/app/sys/model/sys"
	"strconv"

	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppUpdateLogic {
	return &AppUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppUpdateLogic) AppUpdate(in *sysClient.UpdateAppReq) (*sysClient.UpdateAppResp, error) {
	updates := map[string]interface{}{
		"name":                   in.Name,
		"app_id":                 in.AppId,
		"key":                    in.Key,
		"secret":                 in.Secret,
		"grant_type":             in.GrantType,
		"redirect_uri":           in.RedirectUri,
		"access_token_validity":  in.AccessTokenValidity,
		"refresh_token_validity": in.RefreshTokenValidity,
		"additional_information": in.AdditionalInfo,
	}
	var app sys.Client
	result := l.svcCtx.Db.Model(&app).Where("id = ?", in.Id).Updates(updates)
	if result.Error != nil {
		err := errors.New("更新失败:" + result.Error.Error())
		return nil, err
	}

	return &sysClient.UpdateAppResp{
		Data: strconv.FormatInt(result.RowsAffected, 10),
	}, nil
}
