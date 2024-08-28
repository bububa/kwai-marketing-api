package preput

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool/preput"
)

// PredicationTaskCreate 创建投前预估任务
func PredicationTaskCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *preput.PredicationTaskCreateRequest) (*preput.RealTaskResult, error) {
	var resp preput.RealTaskResult
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
