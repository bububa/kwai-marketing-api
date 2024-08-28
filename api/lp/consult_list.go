package lp

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/lp"
)

// ConsultList 获取可用咨询组件列表
func ConsultList(ctx context.Context, clt *core.SDKClient, accessToken string, req *lp.ConsultListRequest) (*lp.ConsultListResponse, error) {
	var resp lp.ConsultListResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
