package app

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/app"
)

// Detail 【应用中心】获取应用详情
func Detail(clt *core.SDKClient, accessToken string, req *app.DetailRequest) (*app.App, error) {
	var resp app.App
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
