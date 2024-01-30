package creative

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/creative"
)

// BatchUpdate 批量修改自定义创意
func BatchUpdate(clt *core.SDKClient, accessToken string, req *creative.BatchUpdateRequest) (*creative.BatchUpdateResponse, error) {
	var resp creative.BatchUpdateResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
