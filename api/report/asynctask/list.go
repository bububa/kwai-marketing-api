package asynctask

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/report/asynctask"
)

func List(ctx context.Context, clt *core.SDKClient, accessToken string, req *asynctask.ListRequest) (*asynctask.ListResponse, error) {
	var resp asynctask.ListResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
