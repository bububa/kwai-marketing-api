package native

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/native"
)

func Upload(ctx context.Context, clt *core.SDKClient, accessToken string, req *native.UploadRequest) (*native.UploadResponse, error) {
	var resp native.UploadResponse
	err := clt.Upload(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
