package dmp

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dmp"
)

// PopulationList 人群列表查询接口
func PopulationList(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.PopulationListRequest) (*dmp.PopulationListResponse, error) {
	var resp *dmp.PopulationListResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
