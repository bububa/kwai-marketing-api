package file

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
)

// AdVideoUploadV2 上传视频v2接口
func AdVideoUploadV2(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdVideoUploadRequestV2) (*file.Video, error) {
	var resp file.Video
	err := clt.Upload(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
