package creative

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/creative"
)

// ActionBarTextList 获取行动号召按钮
func ActionBarTextList(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.ActionBarTextListRequest) ([]string, error) {
	var resp creative.ActionBarTextListResponse
	err := clt.Get(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.ActionBarText, nil
}
