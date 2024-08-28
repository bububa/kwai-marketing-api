package advertiser

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/advertiser"
)

// WhiteList 获取可选白名单接口
func WhiteList(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID uint64) (*advertiser.WhiteListResponse, error) {
	req := &advertiser.WhiteListRequest{
		AdvertiserID: advertiserID,
	}
	var resp advertiser.WhiteListResponse
	err := clt.Get(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
