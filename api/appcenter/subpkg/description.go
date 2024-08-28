package subpkg

import (
	"context"
	"errors"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/subpkg"
)

// Description 【应用中心】修改应用分包备注
func Description(ctx context.Context, clt *core.SDKClient, accessToken string, req *subpkg.DescriptionRequest) error {
	var resp subpkg.ModResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return err
	}
	if !resp.Result {
		return errors.New("修改应用分包备注失败")
	}
	return nil
}
