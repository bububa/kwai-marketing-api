package tool

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool"
)

// TargetingTagsList 获取可选的定向标签
func TargetingTagsList(ctx context.Context, clt *core.SDKClient, accessToken string, req *tool.TargetingTagsListRequest) (*tool.TargetingTag, error) {
	var resp tool.TargetingTag
	err := clt.Get(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
