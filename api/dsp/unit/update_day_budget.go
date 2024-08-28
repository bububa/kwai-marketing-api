package unit

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/unit"
)

// UpdateDayBudget 修改广告组预算
func UpdateDayBudget(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.UpdateDayBudgetRequest) error {
	return clt.Post(ctx, accessToken, req, nil)
}
