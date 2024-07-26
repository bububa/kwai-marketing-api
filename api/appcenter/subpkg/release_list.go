package subpkg

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/subpkg"
)

// ReleaseList 获取新版分包发布列表【单元创编】
func ReleaseList(clt *core.SDKClient, accessToken string, req *subpkg.ReleaseListRequest) (*subpkg.ListResponse, error) {
	var resp subpkg.ListResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
