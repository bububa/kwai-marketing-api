package advertiser

import (
	"github.com/Shinku-Chen/kwai-marketing-api/core"
	"github.com/Shinku-Chen/kwai-marketing-api/model/advertiser"
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
