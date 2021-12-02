package file

import (
	"git.gametaptap.com/tapad/github/kwai-marketing-api/core"
	"git.gametaptap.com/tapad/github/kwai-marketing-api/model/file"
)

// AdVideoUploadV1 上传视频v1接口
func AdVideoUploadV1(clt *core.SDKClient, accessToken string, req *file.AdVideoUploadRequestV1) (string, error) {
	var resp file.Video
	err := clt.Upload(accessToken, req, &resp)
	if err != nil {
		return "", err
	}
	return resp.PhotoID, nil
}
