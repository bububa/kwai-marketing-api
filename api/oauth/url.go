package oauth

import (
	"github.com/Shinku-Chen/kwai-marketing-api/core"
	"github.com/Shinku-Chen/kwai-marketing-api/model/oauth"
)

// Url 生成授权链接
func Url(clt *core.SDKClient, req *oauth.UrlRequest) string {
	req.AppID = clt.AppID()
	return clt.GetUrl(req)
}
