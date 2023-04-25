package creative

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/v2.2/creative"
)

// List 查询广告创意
func AdvancedCreativeList(clt *core.SDKClient, accessToken string, req *creative.AdvancedCreativeListRequest) (*creative.ListAdvancedCreativeResponse, error) {
	var resp creative.ListAdvancedCreativeResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
