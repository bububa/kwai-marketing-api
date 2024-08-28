package app

import (
	"context"
	"errors"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/app"
)

// Online 【应用中心】应用上架
func Online(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.OnlineRequest) error {
	var resp app.OnlineOfflineResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return err
	}
	if !resp.Result {
		return errors.New("应用上架失败")
	}
	return nil
}
