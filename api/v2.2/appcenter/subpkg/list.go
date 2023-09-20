package subpkg

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/v2.2/appcenter/subpkg"
)

func List(clt *core.SDKClient, accessToken string, req *subpkg.ListRequest) (*subpkg.ListResponse, error) {
	var resp subpkg.ListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
