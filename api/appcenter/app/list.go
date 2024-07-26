package app

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/app"
)

// List 【应用中心】获取应用列表
func List(clt *core.SDKClient, accessToken string, req *app.ListRequest) (*app.ListResponse, error) {
	var resp app.ListResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
