package tool

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool"
)

// MAPI Quota 白盒化 https://developers.e.kuaishou.com/docs?docType=DSP&documentId=2244
func QuotaInfo(ctx context.Context, clt *core.SDKClient, accessToken string, req *tool.QuotaInfoRequest) (*tool.QuotaInfoResponse, error) {
	var resp tool.QuotaInfoResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
