package file

import (
	"github.com/Shinku-Chen/kwai-marketing-api/core"
	"github.com/Shinku-Chen/kwai-marketing-api/model/file"
)

// AdAppList 获取应用列表
func AdAppList(clt *core.SDKClient, accessToken string, req *file.AdAppListRequest) (*file.AdAppListResponse, error) {
	var resp file.AdAppListResponse
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
