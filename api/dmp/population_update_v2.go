package dmp

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dmp"
)

// PopulationUpdate 人群包更新接口
func PopulationUpdateV2(clt *core.SDKClient, accessToken string, req *dmp.PopulationUpdateRequestv2) (*dmp.PopulationV2, error) {
	var resp dmp.PopulationV2
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
