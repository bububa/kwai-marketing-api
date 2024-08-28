package dmp

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dmp"
)

// PopulationUpdate 人群包更新接口
func PopulationUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.PopulationUpdateRequest) (*dmp.Population, error) {
	var resp dmp.Population
	err := clt.Upload(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
