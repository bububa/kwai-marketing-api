package comment

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/comment"
)

// SetTop 评论置顶
func SetTop(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.SetTopRequest) (uint64, error) {
	var resp comment.SetTopResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return 0, err
	}
	return resp.CommentID, nil
}
