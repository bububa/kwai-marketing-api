package comment

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/comment"
)

// Shield 屏蔽评论
func Shield(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.ShieldRequest) error {
	return clt.Post(ctx, accessToken, req, nil)
}
