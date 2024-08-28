package target

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/target"
)

// OptionDistanceList 根据店铺名称查询商圈信息
func OptionDistanceList(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.OptionDistanceListRequest) (*target.OptionDistanceListResponse, error) {
	var resp target.OptionDistanceListResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
