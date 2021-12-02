package region

import (
	"git.gametaptap.com/tapad/github/kwai-marketing-api/core"
	"git.gametaptap.com/tapad/github/kwai-marketing-api/model/region"
)

// DistrictList 获取商圈地域定向
func DistrictList(clt *core.SDKClient, accessToken string, advertiserID int64) (map[string]region.District, error) {
	req := &region.DistrictListRequest{
		AdvertiserID: advertiserID,
	}
	resp := make(map[string]region.District)
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
