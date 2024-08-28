package comment

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/comment"
)

// Reply 回复评论
func Reply(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.ReplyRequest) ([]comment.ReplyResult, error) {
	var resp comment.ReplyResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.ReplyResultList, nil
}
