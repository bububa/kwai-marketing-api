package advertiser

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/advertiser"
)

// FundDailyFlows 获取广告账户流水信息
func FundDailyFlows(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.FundDailyFlowsRequest) (*advertiser.FundDailyFlowsResponse, error) {
	var resp advertiser.FundDailyFlowsResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
