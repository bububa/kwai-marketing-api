package region

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/region"
)

// List 获取地域定向
func List(ctx context.Context, clt *core.SDKClient, accessToken string) (map[string]region.Region, error) {
	req := &region.ListRequest{}
	resp := make(map[string]region.Region)
	err := clt.Get(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
