package comment

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/comment"
)

// CancelTop 取消置顶
func CancelTop(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.CancelTopRequest) (uint64, error) {
	var resp comment.CancelTopResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return 0, err
	}
	return resp.CommentID, nil
}
