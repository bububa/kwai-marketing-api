package app

import (
	"errors"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/app"
)

// IosUpdate iOS 应用上报更新
func IosUpdate(clt *core.SDKClient, accessToken string, req *app.IosUpdateRequest) error {
	var resp app.IosUpdateResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return err
	}
	if !resp.Result {
		return errors.New("iOS 应用上报更新失败")
	}
	return nil
}
