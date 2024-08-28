package advertiser

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/advertiser"
)

// BudgetGet 账户日预算查询
func BudgetGet(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID uint64) (*advertiser.Budget, error) {
	req := &advertiser.BudgetGetRequest{
		AdvertiserID: advertiserID,
	}
	var resp advertiser.Budget
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
