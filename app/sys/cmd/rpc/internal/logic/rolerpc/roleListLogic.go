package rolerpclogic

import (
	"context"
	"storm-server/app/sys/model/sys"

	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleListLogic) RoleList(in *sysClient.ListRoleReq) (*sysClient.ListRoleResp, error) {
	var roles []sys.Role
	var total int64
	if err := l.svcCtx.Db.Model(&sys.Role{}).Count(&total).Offset(int((in.Current - 1) * in.PageSize)).Limit(int(in.PageSize)).Find(&roles).Error; err != nil {
		return nil, err
	}
	var result []*sysClient.RoleInfoResp
	for _, item := range roles {
		result = append(result, &sysClient.RoleInfoResp{
			Id:       item.Id,
			Name:     item.Name,
			Code:     item.Code,
			Sort:     int64(item.Sort),
			Remark:   item.Remark,
			Type:     int64(item.Type),
			CreateAt: item.CreateAt.Format("2006-01-02 15:04:05"),
			UpdateAt: item.UpdateAt.Format("2006-01-02 15:04:05"),
		})
	}
	totalPages := (total + in.PageSize - 1) / in.PageSize // 使用向上取整的除法计算总页数

	return &sysClient.ListRoleResp{
		Current:   in.Current,
		PageSize:  in.PageSize,
		List:      result,
		Total:     total,
		TotalPage: totalPages,
	}, nil
}
