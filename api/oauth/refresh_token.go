package oauth

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/oauth"
)

// RefreshToken 刷新tonken
// 请求快手服务器，刷新access_token和refresh_token及token过期时间。
func RefreshToken(ctx context.Context, clt *core.SDKClient, refreshToken string) (*oauth.AccessTokenResponse, error) {
	req := &oauth.RefreshTokenRequest{
		AppID:        clt.AppID(),
		Secret:       clt.Secret(),
		RefreshToken: refreshToken,
	}
	var resp oauth.AccessTokenResponse
	err := clt.Post(ctx, "", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
