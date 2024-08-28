package creative

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/creative"
)

// Update 修改创意
// 【注】联盟广告不支持便利贴图片素材。
func Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.UpdateRequest) (uint64, error) {
	var resp creative.UpdateResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.CreativeID, nil
}
