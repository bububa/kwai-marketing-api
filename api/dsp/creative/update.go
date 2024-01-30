package creative

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/creative"
)

// Update 修改自定义创意
func Update(clt *core.SDKClient, accessToken string, req *creative.UpdateRequest) (uint64, error) {
	var resp creative.UpdateResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return 0, err
	}
	return resp.CreativeID, nil
}
