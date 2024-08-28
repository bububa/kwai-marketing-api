package target

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/dsp/target"
)

// TemplateRelatedUnitList 查询模板关联的广告列表接口
func TemplateRelatedUnitList(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.TemplateRelatedUnitListRequest) (*target.TemplateRelatedUnitListResponse, error) {
	var resp target.TemplateRelatedUnitListResponse
	if err := clt.Post(ctx, accessToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
