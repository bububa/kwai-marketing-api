package campaign

import (
	"github.com/Shinku-Chen/kwai-marketing-api/core"
	"github.com/Shinku-Chen/kwai-marketing-api/model/campaign"
)

// UpdateBudget 修改广告计划预算
func UpdateBudget(clt *core.SDKClient, accessToken string, req *campaign.UpdateBudgetRequest) error {
	return clt.Post(accessToken, req, nil)
}
