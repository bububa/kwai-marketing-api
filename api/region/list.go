package region

import (
	"github.com/Shinku-Chen/kwai-marketing-api/core"
	"github.com/Shinku-Chen/kwai-marketing-api/model/region"
)

// List 获取地域定向
func List(clt *core.SDKClient, accessToken string) (map[string]region.Region, error) {
	req := &region.ListRequest{}
	resp := make(map[string]region.Region)
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
