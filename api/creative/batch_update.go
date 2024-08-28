package creative

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/creative"
)

// BatchUpdate 批量创建&修改创意
// 注：该接口可实现创意的批量创建&更新，上传creative_id为更新，不上传creative_id则为创建。
func BatchUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.BatchUpdateRequest) (*creative.BatchUpdateResponse, error) {
	var resp creative.BatchUpdateResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
