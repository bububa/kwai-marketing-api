package oauth

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/oauth"
)

// Url 生成授权链接
func Url(clt *core.SDKClient, req *oauth.UrlRequest) string {
	req.AppID = clt.AppID()
	return clt.GetUrl(req)
}
