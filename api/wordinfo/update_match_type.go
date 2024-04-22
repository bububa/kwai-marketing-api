package wordinfo

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/wordinfo"
)

// UpdateMatchType 修改关键词匹配方式
func UpdateMatchType(clt *core.SDKClient, accessToken string, req *wordinfo.UpdateMatchTypeRequest) ([]uint64, error) {
	var resp wordinfo.UpdateMatchTypeResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return resp.WordInfoIDs, nil
}
