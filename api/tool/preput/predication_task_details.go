package preput

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool/preput"
)

// PredicationTaskDetails 投前预估详情
func PredicationTaskDetails(ctx context.Context, clt *core.SDKClient, accessToken string, req *preput.PredicationTaskCreateRequest) (*preput.AdPredicationTaskDetail, error) {
	var resp preput.AdPredicationTaskDetail
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
