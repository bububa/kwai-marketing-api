package file

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
)

// AdVideoList 查询视频信息list接口
func VideoListByCursor(clt *core.SDKClient, accessToken string, req *file.VideoListByCursorRequest) (*file.VideoListByCursorResponse, error) {
	var resp file.VideoListByCursorResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

