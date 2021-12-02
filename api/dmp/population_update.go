package dmp

import (
	"git.gametaptap.com/tapad/github/kwai-marketing-api/core"
	"git.gametaptap.com/tapad/github/kwai-marketing-api/model/dmp"
)

// PopulationUpdate 人群包更新接口
func PopulationUpdate(clt *core.SDKClient, accessToken string, req *dmp.PopulationUpdateRequest) (*dmp.Population, error) {
	var resp dmp.Population
	err := clt.Upload(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
