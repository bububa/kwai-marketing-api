package creative

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/creative"
)

// CreativeTagAdvise 创意标签填写建议
func CreativeTagAdvise(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.CreativeTagAdviseRequest) (*creative.CreativeTagAdviseResponse, error) {
	var resp creative.CreativeTagAdviseResponse
	err := clt.Get(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
