package userrpclogic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"storm-server/app/sys/model/sys"
	"storm-server/common/tool"
	"time"

	"storm-server/app/sys/cmd/rpc/internal/svc"
	"storm-server/app/sys/cmd/rpc/sysClient"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *sysClient.LoginReq) (*sysClient.LoginResp, error) {
	//1.用户查询
	var user sys.User
	err := l.svcCtx.Db.Take(&user, "account = ?", in.Account).Error
	if err != nil {
		err = errors.New("用户名或密码错误")
		return nil, errors.Wrap(err, "LoginLogic.Login")
	}

	//2.校验密码
	if !tool.CheckPwd(user.Password, in.Password) {
		err = errors.New("用户名或密码错误")
		return nil, errors.Wrap(err, "LoginLogic.Login")
	}
	//3.查询app 根据app获取token有效期
	var client sys.Client
	tx := l.svcCtx.Db.Take(&client, "app_id = ?", in.AppId)
	if tx.Error != nil {
		err = errors.New("app不存在")
		return nil, errors.Wrap(err, "LoginLogic.Login")
	}

	now := time.Now().Unix()
	accessSecret := l.svcCtx.Config.JwtAuth.AccessSecret
	validity := client.AccessTokenValidity
	userId := user.Id
	account := user.Account
	token, err := l.getJwtToken(accessSecret, now, validity, userId, account)
	if err != nil {
		return nil, err
	}

	perms := []string{"/admin", "/dashboard"}
	return &sysClient.LoginResp{
		AccessToken:  token,
		AccessExpire: validity,
		RefreshAfter: 0,
		UserInfo: &sysClient.UserInfo{
			Id:       user.Id,
			Name:     user.Nickname,
			RealName: user.Name,
			Account:  user.Account,
			Phone:    user.Phone,
		},
		Permissions: perms,
	}, nil
}

// 生成jwt的token
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds int64, userId string, account string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	claims["account"] = account
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
