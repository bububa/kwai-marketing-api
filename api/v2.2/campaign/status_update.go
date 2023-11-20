package campaign


import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/v2.2/campaign"
	"github.com/bububa/kwai-marketing-api/model/v2.2/unit"
)

// Create 创建广告组
func StatusUpdate(clt *core.SDKClient, accessToken string, req *unit.StatusUpdateRequest) ([]uint64, error) {
	var resp campaign.StatusUpdateResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.CampaignIds, nil
}
