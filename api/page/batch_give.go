package page

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/page"
)

// BatchGive 批量转赠
func BatchGive(ctx context.Context, clt *core.SDKClient, accessToken string, req *page.BatchGiveRequest) error {
	var resp page.BatchGiveResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return err
	}
	if resp.IsError() {
		return resp
	}
	return nil
}
