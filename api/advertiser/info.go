package advertiser

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/advertiser"
)

// 获取广告账户信息
func Info(clt *core.SDKClient, accessToken string, advertiserID int64) (*advertiser.Info, error) {
	req := &advertiser.InfoRequest{
		AdvertiserID: advertiserID,
	}
	var resp advertiser.Info
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
