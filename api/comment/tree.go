package comment

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/comment"
)

// Tree 评论树查询
func Tree(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.TreeRequest) (*comment.TreeResponse, error) {
	var resp comment.TreeResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
