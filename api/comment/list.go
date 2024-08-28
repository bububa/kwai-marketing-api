package comment

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/comment"
)

// List 评论列表数据查询
func List(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.ListRequest) (*comment.ListResponse, error) {
	var resp comment.ListResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
