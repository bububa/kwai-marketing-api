package asset

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/asset"
)

// AdvCardCreate 创建高级创意接口
func AdvCardCreate(clt *core.SDKClient, accessToken string, req *asset.AdvCardCreateRequest) ([]int64, error) {
	var resp asset.AdvCardCreateResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.AdvList, nil
}
