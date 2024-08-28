package wordinfo

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/wordinfo"
)

// Create 创建关键词
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *wordinfo.CreateRequest) (*wordinfo.CreateResponse, error) {
	var resp wordinfo.CreateResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
