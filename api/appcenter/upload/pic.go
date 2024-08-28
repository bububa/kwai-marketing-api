package upload

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/upload"
)

// Pic 上传图片
func Pic(ctx context.Context, clt *core.SDKClient, accessToken string, req *upload.PicRequest) (string, error) {
	var resp upload.PicResponse
	if err := clt.Upload(ctx, accessToken, req, &resp); err != nil {
		return "", err
	}
	return resp.URL, nil
}
