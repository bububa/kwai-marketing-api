package subpkg

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/v2.2/appcenter/subpkg"
)

func Create(clt *core.SDKClient, accessToken string, req *subpkg.CreateRequest) (*subpkg.CreateResponse, error) {
	var resp subpkg.CreateResponse
	err := clt.Post(accessToken, req, &resp.Item)
	if err != nil {
		return &resp, err
	}
	return &resp, nil
}
