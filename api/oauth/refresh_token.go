package oauth

import (
	"git.gametaptap.com/tapad/github/kwai-marketing-api/core"
	"git.gametaptap.com/tapad/github/kwai-marketing-api/model/oauth"
)

// RefreshToken 刷新tonken
// 请求快手服务器，刷新access_token和refresh_token及token过期时间。
func RefreshToken(clt *core.SDKClient, refreshToken string) (*oauth.AccessTokenResponse, error) {
	req := &oauth.RefreshTokenRequest{
		AppID:        clt.AppID(),
		Secret:       clt.Secret(),
		RefreshToken: refreshToken,
	}
	var resp oauth.AccessTokenResponse
	err := clt.Post("", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
