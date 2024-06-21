package oauth

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/oauth"
)

// ApprovalList 拉取token下授权广告账户接口
func ApprovalList(clt *core.SDKClient, accessToken string) ([]uint64, error) {
	req := &oauth.ApprovalListRequest{
		AppID:       clt.AppID(),
		Secret:      clt.Secret(),
		AccessToken: accessToken,
	}
	var resp oauth.ApprovalListResponse
	err := clt.Post("", req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Details, nil
}
