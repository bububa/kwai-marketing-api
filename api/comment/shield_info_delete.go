package comment

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/comment"
)

// ShieldInfoDelete 删除屏蔽评论信息
func ShieldInfoDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.ShieldInfoDeleteRequest) error {
	return clt.Post(ctx, accessToken, req, nil)
}
