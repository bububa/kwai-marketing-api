package page

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/page"
)

// BatchGive 批量转赠
func BatchGive(clt *core.SDKClient, accessToken string, req *page.BatchGiveRequest) error {
	var resp page.BatchGiveResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return err
	}
	if resp.IsError() {
		return resp
	}
	return nil
}
