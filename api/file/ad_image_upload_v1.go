package file

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
)

// AdImageUploadV1 上传图片v1接口
func AdImageUploadV1(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdImageUploadRequestV1) (*file.Image, error) {
	var resp file.Image
	err := clt.Upload(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
