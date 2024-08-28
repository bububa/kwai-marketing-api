package native

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/native"
)

func PhotoList(ctx context.Context, clt *core.SDKClient, accessToken string, req *native.PhotoListRequest) (*native.PhotoListResponse, error) {
	var resp native.PhotoListResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
