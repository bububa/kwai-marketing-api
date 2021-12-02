package oauth

import (
	"git.gametaptap.com/tapad/github/kwai-marketing-api/core"
	"git.gametaptap.com/tapad/github/kwai-marketing-api/model/oauth"
)

// Url 生成授权链接
func Url(clt *core.SDKClient, req *oauth.UrlRequest) string {
	req.AppID = clt.AppID()
	return clt.GetUrl(req)
}
