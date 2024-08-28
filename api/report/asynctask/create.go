package asynctask

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/report/asynctask"
)

// Create 创建历史数据查询任务
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *asynctask.CreateRequest) (*asynctask.CreateResponse, error) {
	var resp asynctask.CreateResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
