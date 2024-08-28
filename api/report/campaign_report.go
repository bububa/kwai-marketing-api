package report

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/report"
)

// CampaignReport 广告计划数据（实时）
func CampaignReport(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.CampaignReportRequest) (*report.ReportResponse, error) {
	var resp report.ReportResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
