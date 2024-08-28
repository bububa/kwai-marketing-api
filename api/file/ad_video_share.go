package file

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
)

// AdVideoShare 视频库-推送视频
func AdVideoShare(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdVideoShareRequest) ([]file.AdVideoShareDetail, error) {
	var resp file.AdVideoShareResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Details, nil
}
