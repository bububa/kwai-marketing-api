package subpkg

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/v2.2/appcenter/subpkg"
)

func Update(clt *core.SDKClient, accessToken string, req *subpkg.UpdateRequest) (*subpkg.UpdateResponse, error) {
	var resp subpkg.UpdateResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return &resp, err
	}
	return &resp, nil
}