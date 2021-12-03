package file

import (
	"github.com/Shinku-Chen/kwai-marketing-api/core"
	"github.com/Shinku-Chen/kwai-marketing-api/model/file"
)

// AdImageUploadV1 上传图片v1接口
func AdImageUploadV1(clt *core.SDKClient, accessToken string, req *file.AdImageUploadRequestV1) (*file.Image, error) {
	var resp file.Image
	err := clt.Upload(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
