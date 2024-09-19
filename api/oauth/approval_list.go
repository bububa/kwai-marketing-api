package oauth

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/oauth"
)

// ApprovalList 拉取token下授权广告账户接口
func ApprovalList(ctx context.Context, clt *core.SDKClient, accessToken string, pageNo int, pageSize int) (*oauth.ApprovalListResponse, error) {
	req := &oauth.ApprovalListRequest{
		AppID:       clt.AppID(),
		Secret:      clt.Secret(),
		AccessToken: accessToken,
		PageNo:      pageNo,
		PageSize:    pageSize,
	}
	var resp oauth.ApprovalListResponse
	if err := clt.Post(ctx, "", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
