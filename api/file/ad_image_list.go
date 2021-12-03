package file

import (
	"github.com/Shinku-Chen/kwai-marketing-api/core"
	"github.com/Shinku-Chen/kwai-marketing-api/model/file"
)

// AdImageList 查询图片信息list接口
func AdImageList(clt *core.SDKClient, accessToken string, req *file.AdImageListRequest) (*file.AdImageListResponse, error) {
	var resp file.AdImageListResponse
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
