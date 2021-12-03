package campaign

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/campaign"
)

// UpdateBudget 修改广告计划预算
func UpdateBudget(clt *core.SDKClient, accessToken string, req *campaign.UpdateBudgetRequest) error {
	return clt.Post(accessToken, req, nil)
}
