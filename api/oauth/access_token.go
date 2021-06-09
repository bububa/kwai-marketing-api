package oauth

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/oauth"
)

// 获取token
// 利用授权码auth_code，请求快手服务器，获取access_token和refresh_token及当前账户的广告主ID。
func AccessToken(clt *core.SDKClient, authCode string) (*oauth.AccessTokenResponse, error) {
	req := &oauth.AccessTokenRequest{
		AppID:    clt.AppID(),
		Secret:   clt.Secret(),
		AuthCode: authCode,
	}
	var resp oauth.AccessTokenResponse
	err := clt.Post("", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
