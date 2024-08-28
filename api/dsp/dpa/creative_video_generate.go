package dpa

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/dpa"
)

// CreativeVideoGenerate 批量模板合成SDPA创意视频
func CreativeVideoGenerate(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.CreativeVideoGenerateRequest) ([]dpa.GenerateVideoResult, error) {
	var resp dpa.CreativeVideoGenerateResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return resp.VideoInfos, nil
}
