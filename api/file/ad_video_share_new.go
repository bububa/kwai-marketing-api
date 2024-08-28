package file

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
)

// AdVideoShareNew 视频库-推送视频(新版)
func AdVideoShareNew(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdVideoShareNewRequest) (*file.AdVideoShareNewResponse, error) {
	var resp file.AdVideoShareNewResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
