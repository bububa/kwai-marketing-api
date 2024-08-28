package unit

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/unit"
)

// GetMonitorURLs 批量获取监测链接接口
func GetMonitorURLs(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.GetMonitorURLsRequest) ([]unit.UnitMonitorURL, error) {
	var resp unit.GetMonitorURLsResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return resp.UnitMonitorURLs, nil
}
