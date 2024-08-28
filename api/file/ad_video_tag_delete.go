package file

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
)

// AdVideoTagDelete 视频库-删除视频标签
func AdVideoTagDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdVideoTagDeleteRequest) error {
	return clt.Post(ctx, accessToken, req, nil)
}
