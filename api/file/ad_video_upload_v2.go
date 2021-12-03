package file

import (
	"github.com/Shinku-Chen/kwai-marketing-api/core"
	"github.com/Shinku-Chen/kwai-marketing-api/model/file"
)

// AdVideoUploadV2 上传视频v2接口
func AdVideoUploadV2(clt *core.SDKClient, accessToken string, req *file.AdVideoUploadRequestV2) (*file.Video, error) {
	var resp file.Video
	err := clt.Upload(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
