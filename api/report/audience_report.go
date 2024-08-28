package report

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/report"
)

// AudienceReport 人群分析报表
func AudienceReport(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.AudienceReportRequest) (*report.ReportResponse, error) {
	var resp report.ReportResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
