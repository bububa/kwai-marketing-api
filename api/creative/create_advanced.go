package creative

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/creative"
)

// Create 创建创意
// 【注】联盟广告不支持便利贴图片素材，只有联盟广告支持横版竖版图片素材。
func CreateAdvanced(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.CreateAdvancedRequest) (uint64, error) {
	var resp creative.CreateAdvancedResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.UnitId, nil
}
