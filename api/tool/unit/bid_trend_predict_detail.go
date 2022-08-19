package unit

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool/unit"
)

// BidTrendPredictDetail 获取广告组投放预估曲线
func BidTrendPredictDetail(clt *core.SDKClient, accessToken string, req *unit.BidTrendPredictDetailRequest) (*unit.BidTrendPredict, error) {
	var resp unit.BidTrendPredict
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
