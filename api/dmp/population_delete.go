package dmp

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dmp"
)

// PopulationDelete 人群包删除接口
func PopulationDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.PopulationDeleteRequest) error {
	return clt.Post(ctx, accessToken, req, nil)
}
