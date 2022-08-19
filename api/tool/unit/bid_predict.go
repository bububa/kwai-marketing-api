package unit

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool/unit"
)

// BidPredict 获取广告组曝光、转化预估
func BidPredict(clt *core.SDKClient, accessToken string, req *unit.BidPredictRequest) (*unit.BidPredict, error) {
	var resp unit.BidPredict
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
