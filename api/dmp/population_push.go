package dmp

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dmp"
)

// PopulationPush 人群包上线接口
func PopulationPush(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.PopulationPushRequest) error {
	return clt.Post(ctx, accessToken, req, nil)
}
