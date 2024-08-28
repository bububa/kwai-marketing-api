package comment

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/comment"
)

// ShieldInfoList 获取屏蔽评论信息列表
func ShieldInfoList(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.ShieldInfoListRequest) (*comment.ShieldInfoListResponse, error) {
	var resp comment.ShieldInfoListResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
