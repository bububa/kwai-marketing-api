package advertiser

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/advertiser"
)

// Info 获取广告账户信息
func Info(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID uint64) (*advertiser.Info, error) {
	req := &advertiser.InfoRequest{
		AdvertiserID: advertiserID,
	}
	var resp advertiser.Info
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
