package preput

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool/preput"
)

// PredicationTaskList 投前预估列表页接口
func PredicationTaskList(ctx context.Context, clt *core.SDKClient, accessToken string, req *preput.PredicationTaskListRequest) (*preput.PredicationTaskListResponse, error) {
	var resp preput.PredicationTaskListResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
