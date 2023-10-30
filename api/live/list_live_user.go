package live

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/live"
)

// ListLiveUser 获取主播列表
func ListLiveUser(clt *core.SDKClient, accessToken string, req *live.ListLiveUserRequest) (*live.ListLiveUserResponse, error) {
	var resp live.ListLiveUserResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
