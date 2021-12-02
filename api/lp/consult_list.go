package lp

import (
	"git.gametaptap.com/tapad/github/kwai-marketing-api/core"
	"git.gametaptap.com/tapad/github/kwai-marketing-api/model/lp"
)

// ConsultList 获取可用咨询组件列表
func ConsultList(clt *core.SDKClient, accessToken string, req *lp.ConsultListRequest) (*lp.ConsultListResponse, error) {
	var resp lp.ConsultListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
