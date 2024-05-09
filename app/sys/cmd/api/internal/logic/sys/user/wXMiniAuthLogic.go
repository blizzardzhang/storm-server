package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"storm-server/common/xerr"

	"storm-server/app/sys/cmd/api/internal/svc"
	"storm-server/app/sys/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrWxMiniAuthFailError = xerr.NewErrMsg("wechat mini auth fail")

const (
	code2SessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	checkEncryptedDataURL = "https://api.weixin.qq.com/wxa/business/checkencryptedmsg?access_token=%s"
)

type ResCode2Session struct {
	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足UnionID下发条件的情况下会返回
}

type WXMiniAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWXMiniAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WXMiniAuthLogic {
	return &WXMiniAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WXMiniAuthLogic) WXMiniAuth(req *types.WXMiniAuthReq) (resp *types.WXMiniAuthResp, err error) {
	//1.发起微信请求获取openid
	appId := l.svcCtx.Config.WxMiniConf.AppId
	secret := l.svcCtx.Config.WxMiniConf.Secret
	wxResult, err := http.Get(fmt.Sprintln(code2SessionURL, appId, secret, req.Code))
	if err != nil {
		return nil, errors.Wrapf(ErrWxMiniAuthFailError, "发起授权请求失败 err : %v , code : %s  , authResult : %+v", err, req.Code, wxResult)
	}
	defer wxResult.Body.Close()
	if wxResult.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", code2SessionURL, wxResult.StatusCode)
	}
	//2.解析微信返回结果
	var result ResCode2Session
	all, err := io.ReadAll(wxResult.Body)
	if err != nil {
		return nil, errors.Wrapf(ErrWxMiniAuthFailError, "读取授权结果失败 err : %v , code : %s  , authResult : %+v", err, req.Code, wxResult)
	}
	err = json.Unmarshal(all, &result)

	//3.注册并登录或者登录
	//openID := result.OpenID

	return nil, nil

}
