package unit

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/unit"
)

// UpdateStatus 修改广告组状态
func UpdateStatus(clt *core.SDKClient, accessToken string, req *unit.UpdateStatusRequest) ([]uint64, error) {
	var resp unit.UpdateStatusResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.UnitIDs, nil
}
