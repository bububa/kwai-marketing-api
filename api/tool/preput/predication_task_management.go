package preput

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool/preput"
)

// PredicationTaskManagement 投前预估任务管理接口
func PredicationTaskManagement(ctx context.Context, clt *core.SDKClient, accessToken string, req *preput.PredicationTaskManagementRequest) (*preput.RealTaskResult, error) {
	var resp preput.RealTaskResult
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
