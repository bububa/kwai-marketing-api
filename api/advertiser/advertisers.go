package advertiser

import (
	"git.gametaptap.com/tapad/github/kwai-marketing-api/core"
	"git.gametaptap.com/tapad/github/kwai-marketing-api/model/advertiser"
)

// AdvertisersGet   获取罗盘绑定广告主列
func AdvertisersGet(clt *core.SDKClient, accessToken string, advertiserID int64) (*advertiser.AdvertisersResponse, error) {
	req := &advertiser.AdvertisersRequest{
		AdvertiserID: advertiserID,
	}
	var resp advertiser.AdvertisersResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
