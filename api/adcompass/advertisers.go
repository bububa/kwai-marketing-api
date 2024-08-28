package adcompass

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/adcompass"
)

// Advertisers 获取罗盘绑定广告主列
func Advertisers(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID uint64) ([]adcompass.Advertiser, error) {
	req := &adcompass.AdvertisersRequest{
		AdvertiserID: advertiserID,
	}
	var resp adcompass.AdvertisersResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Details, nil
}
