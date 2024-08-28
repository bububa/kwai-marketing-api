package upload

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/appcenter/upload"
)

// Apk 上传 APK 文件
func Apk(ctx context.Context, clt *core.SDKClient, accessToken string, req *upload.ApkRequest) (string, error) {
	var resp upload.ApkResponse
	if err := clt.Upload(ctx, accessToken, req, &resp); err != nil {
		return "", err
	}
	return resp.BlobStoreKey, nil
}
