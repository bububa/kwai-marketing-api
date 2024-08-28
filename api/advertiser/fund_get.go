package advertiser

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/advertiser"
)

// FundGet 获取广告账户余额信息
func FundGet(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID uint64) (float64, error) {
	req := &advertiser.FundGetRequest{
		AdvertiserID: advertiserID,
	}
	var resp advertiser.FundGetResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.Balance, nil
}
