package dmp

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dmp"
)

// PopulationAccountsPush 人群包跨账户推送
func PopulationAccountsPush(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.PopulationAccountsPushRequest) (*dmp.PopulationAccountsPushResponse, error) {
	var resp dmp.PopulationAccountsPushResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
