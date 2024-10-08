package unit

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/v2.2/unit"
)

// 更新广告组预算
func BudgetUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.BudgetUpdateRequest) ([]uint64, error) {
	var resp unit.BudgetUpdateResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.UnitIds, nil
}
