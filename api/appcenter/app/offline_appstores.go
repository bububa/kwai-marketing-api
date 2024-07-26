package app

import (
	"errors"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/app"
)

// OfflineAppStores 【应用中心】应用商店上下架
func OfflineAppStores(clt *core.SDKClient, accessToken string, req *app.OfflineAppStoresRequest) error {
	var resp app.OnlineOfflineResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return err
	}
	if !resp.Result {
		return errors.New("应用商店上下架失败")
	}
	return nil
}
