package file

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
)

// AdVideoUploadV1 上传视频v1接口
func AdVideoUploadV1(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdVideoUploadRequestV1) (string, error) {
	var resp file.Video
	err := clt.Upload(ctx, accessToken, req, &resp)
	if err != nil {
		return "", err
	}
	return resp.PhotoID, nil
}
