package target

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/target"
)

// TemplateSyncHistory 模板同步失败查询
func TemplateSyncHistory(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.TemplateSyncHistoryRequest) (*target.TemplateSyncHistoryResponse, error) {
	var resp target.TemplateSyncHistoryResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
