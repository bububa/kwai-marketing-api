package dmp

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dmp"
)

// PopulationUpload 人群包上传接口
func PopulationUpload(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.PopulationUploadRequest) (*dmp.Population, error) {
	var resp dmp.Population
	err := clt.Upload(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
