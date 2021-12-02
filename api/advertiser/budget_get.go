package advertiser

import (
	"git.gametaptap.com/tapad/github/kwai-marketing-api/core"
	"git.gametaptap.com/tapad/github/kwai-marketing-api/model/advertiser"
)

// BudgetGet 账户日预算查询
func BudgetGet(clt *core.SDKClient, accessToken string, advertiserID int64) (*advertiser.Budget, error) {
	req := &advertiser.BudgetGetRequest{
		AdvertiserID: advertiserID,
	}
	var resp advertiser.Budget
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
