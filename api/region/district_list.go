package region

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/region"
)

// DistrictList 获取商圈地域定向
func DistrictList(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID uint64) (map[string]region.District, error) {
	req := &region.DistrictListRequest{
		AdvertiserID: advertiserID,
	}
	resp := make(map[string]region.District)
	err := clt.Get(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
