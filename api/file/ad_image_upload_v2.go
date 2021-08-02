package file

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
)

// AdImageUploadV2 上传图片v2接口
func AdImageUploadV2(clt *core.SDKClient, accessToken string, req *file.AdImageUploadRequestV2) (*file.Image, error) {
	var resp file.Image
	err := clt.Upload(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
