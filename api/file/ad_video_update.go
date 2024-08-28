package file

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
)

// AdVideoUpdate 视频库-批量更新视频功能
func AdVideoUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdVideoUpdateRequest) ([]string, error) {
	var ret file.AdVideoUpdateResponse
	if err := clt.Post(ctx, accessToken, req, &ret); err != nil {
		return nil, err
	}
	return ret.PhotoIDs, nil
}
