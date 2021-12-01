package advertiser

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/advertiser"
)

// AdvertisersGet   获取罗盘绑定广告主列
func AdvertisersGet(clt *core.SDKClient, accessToken string, advertiserID int64) (*advertiser.AdvertisersResponse, error) {
	req := &advertiser.AdvertisersRequest{
		AdvertiserID: advertiserID,
	}
	var resp advertiser.AdvertisersResponse
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
