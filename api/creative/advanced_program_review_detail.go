package creative

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/creative"
)

// AdvancedProgramReviewDetail 获取程序化创意2.0审核信息
func AdvancedProgramReviewDetail(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.AdvancedProgramReviewDetailRequest) (*creative.AdvancedProgramReviewDetail, error) {
	var resp creative.AdvancedProgramReviewDetail
	err := clt.Get(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
