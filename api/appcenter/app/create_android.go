package app

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/app"
)

// CreateAndroid 创建Android应用
func CreateAndroid(clt *core.SDKClient, accessToken string, req *app.CreateAndroidRequest) (*app.App, error) {
	var resp app.App
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
