package subpkg

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/subpkg"
)

// Add 【应用中心】新建应用分包
func Add(clt *core.SDKClient, accessToken string, req *subpkg.AddRequest) ([]subpkg.SubPackage, error) {
	var resp []subpkg.SubPackage
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return resp, err
	}
	return resp, nil
}
