package app

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/app"
)

// UpdateAndroid 更新Android应用
func UpdateAndroid(clt *core.SDKClient, accessToken string, req *app.UpdateAndroidRequest) (*app.App, error) {
	var resp app.App
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
