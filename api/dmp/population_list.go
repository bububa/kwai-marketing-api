package dmp

import (
	"github.com/Shinku-Chen/kwai-marketing-api/core"
	"github.com/Shinku-Chen/kwai-marketing-api/model/dmp"
)

// PopulationList 人群列表查询接口
func PopulationList(clt *core.SDKClient, accessToken string, req *dmp.PopulationListRequest) ([]dmp.Population, error) {
	var resp []dmp.Population
	err := clt.Get(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
