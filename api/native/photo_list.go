package native

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/native"
)

func PhotoList(clt *core.SDKClient, accessToken string, req *native.PhotoListRequest) (*native.PhotoListResponse, error) {
	var resp native.PhotoListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
