package unit

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/unit"
)

// UpdateBid 修改广告组出价
func UpdateBid(clt *core.SDKClient, accessToken string, req *unit.UpdateBidRequest) error {
	return clt.Post(accessToken, req, nil)
}
