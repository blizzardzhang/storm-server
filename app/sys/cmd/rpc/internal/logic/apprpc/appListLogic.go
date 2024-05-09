package apprpclogic

import (
	"context"
	"storm-server/app/sys/model/sys"

	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppListLogic {
	return &AppListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppListLogic) AppList(in *sysClient.ListAppReq) (*sysClient.ListAppResp, error) {
	var apps []sys.Client
	var total int64
	//特别注意，需要加 Db.Model(&sys.Client{}).Count(&total) 分页查询才正常
	tx := l.svcCtx.Db.Model(&sys.Client{}).Count(&total).Offset(int((in.Current - 1) * in.PageSize)).Limit(int(in.PageSize)).Find(&apps)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var result []*sysClient.AppInfoResp
	for _, item := range apps {
		result = append(result, &sysClient.AppInfoResp{
			Id:                   item.Id,
			AppId:                item.ClientId,
			Name:                 item.Name,
			Key:                  item.Key,
			Secret:               item.Secret,
			GrantType:            item.GrantType,
			RedirectUri:          item.RedirectUri,
			AdditionalInfo:       item.AdditionalInformation,
			AccessTokenValidity:  item.AccessTokenValidity,
			RefreshTokenValidity: item.RefreshTokenValidity,
			CreateAt:             item.CreateAt.Format("2006-01-02 15:04:05"),
			UpdateAt:             item.UpdateAt.Format("2006-01-02 15:04:05"),
			Status:               int64(item.Status),
		})
	}
	totalPages := (total + in.PageSize - 1) / in.PageSize // 使用向上取整的除法计算总页数

	return &sysClient.ListAppResp{
		Current:   in.Current,
		PageSize:  in.PageSize,
		List:      result,
		Total:     total,
		TotalPage: totalPages,
	}, nil
}
