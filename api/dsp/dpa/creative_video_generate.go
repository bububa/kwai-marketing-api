package dpa

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/dpa"
)

// CreativeVideoGenerate 批量模板合成SDPA创意视频
func CreativeVideoGenerate(clt *core.SDKClient, accessToken string, req *dpa.CreativeVideoGenerateRequest) ([]dpa.GenerateVideoResult, error) {
	var resp dpa.CreativeVideoGenerateResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return resp.VideoInfos, nil
}
