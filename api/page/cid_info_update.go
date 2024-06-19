package page

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/page"
)

// CidInfoUpdate 魔力建站落地页更新CID信息
func CidInfoUpdate(clt *core.SDKClient, accessToken string, req *page.CidInfoUpdateRequest) (uint64, error) {
	var resp page.CidInfoUpdateResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return 0, err
	}
	return resp.PageID.Value(), nil
}
