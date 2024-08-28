package report

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/report"
)

// WordInfoReport 关键词报表
func WordInfoReport(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.WordInfoReportRequest) (*report.ReportResponse, error) {
	var resp report.ReportResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
