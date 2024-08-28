package unit

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/unit"
)

// UpdateBid 修改广告组出价
func UpdateBid(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.UpdateBidRequest) error {
	return clt.Post(ctx, accessToken, req, nil)
}
