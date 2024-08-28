package unit

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/unit"
)

// BatchUpdateMonitorURLs 监测链接批量更新接口
func BatchUpdateMonitorURLs(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.BatchUpdateMonitorURLsRequest) ([]unit.UnitMonitorURL, error) {
	var resp unit.GetMonitorURLsResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return resp.UnitMonitorURLs, nil
}
