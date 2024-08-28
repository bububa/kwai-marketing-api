package unit

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/v2.2/unit"
)

func BidUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.BidUpdateRequest) ([]uint64, error) {
	var resp unit.BidUpdateResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.UnitIds, nil
}
