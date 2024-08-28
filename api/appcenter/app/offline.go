package app

import (
	"context"
	"errors"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/app"
)

// Offline 【应用中心】应用下架
func Offline(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.OfflineRequest) error {
	var resp app.OnlineOfflineResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return err
	}
	if !resp.Result {
		return errors.New("应用下架失败")
	}
	return nil
}
