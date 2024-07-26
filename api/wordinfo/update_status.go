package wordinfo

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/wordinfo"
)

// UpdateStatus 修改关键词投放状态
func UpdateStatus(clt *core.SDKClient, accessToken string, req *wordinfo.UpdateStatusRequest) ([]uint64, error) {
	var resp wordinfo.UpdateStatusResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return resp.WordInfoIDs, nil
}
