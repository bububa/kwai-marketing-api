package app

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/v2.2/appcenter/app"
)

func List(clt *core.SDKClient, accessToken string, req *app.ListRequest) (*app.ListResponse, error) {
	var resp app.ListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

