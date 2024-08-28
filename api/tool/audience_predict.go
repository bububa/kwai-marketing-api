package tool

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool"
)

// AudiencePredict 定向人群预估查询
func AudiencePredict(ctx context.Context, clt *core.SDKClient, accessToken string, req *tool.AudiencePredictRequest) (int64, error) {
	var resp tool.AudiencePredictResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.AudiencePredictNum, nil
}
