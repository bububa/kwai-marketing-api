package advertiser

import (
	"github.com/Shinku-Chen/kwai-marketing-api/core"
	"github.com/Shinku-Chen/kwai-marketing-api/model/advertiser"
)

// FundDailyFlows 获取广告账户流水信息
func FundDailyFlows(clt *core.SDKClient, accessToken string, req *advertiser.FundDailyFlowsRequest) (*advertiser.FundDailyFlowsResponse, error) {
	var resp advertiser.FundDailyFlowsResponse
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
