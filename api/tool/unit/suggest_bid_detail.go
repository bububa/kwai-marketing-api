package unit

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool/unit"
)

// SuggestBidDetail 获取广告组出价建议
func SuggestBidDetail(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.SuggestBidDetailRequest) ([]unit.SuggestBidUnit, error) {
	var resp []unit.SuggestBidUnit
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
