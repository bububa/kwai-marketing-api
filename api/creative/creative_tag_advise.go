package creative

import (
	"git.gametaptap.com/tapad/github/kwai-marketing-api/core"
	"git.gametaptap.com/tapad/github/kwai-marketing-api/model/creative"
)

// CreativeTagAdvise 创意标签填写建议
func CreativeTagAdvise(clt *core.SDKClient, accessToken string, req *creative.CreativeTagAdviseRequest) (*creative.CreativeTagAdviseResponse, error) {
	var resp creative.CreativeTagAdviseResponse
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
