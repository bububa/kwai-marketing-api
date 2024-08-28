package target

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/target"
)

// TemplateList 查询定向模板接口
func TemplateList(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.TemplateListRequest) (*target.TemplateListResponse, error) {
	var resp target.TemplateListResponse
	err := clt.Get(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
