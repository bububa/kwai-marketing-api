package tool

import (
	"git.gametaptap.com/tapad/github/kwai-marketing-api/core"
	"git.gametaptap.com/tapad/github/kwai-marketing-api/model/tool"
)

// AudiencePredict 定向人群预估查询
func AudiencePredict(clt *core.SDKClient, accessToken string, req *tool.AudiencePredictRequest) (int64, error) {
	var resp tool.AudiencePredictResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.AudiencePredictNum, nil
}
