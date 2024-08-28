package dmp

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dmp"
)

// PopulationUpload 人群包上传接口
func PopulationUploadFileV2(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.PopulationUploadFileRequest) (*dmp.FileV2, error) {
	var resp dmp.FileV2
	err := clt.Upload(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
