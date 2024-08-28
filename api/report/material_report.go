package report

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/report"
)

// MaterialReport 素材报表
func MaterialReport(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.MaterialReportRequest) (*report.ReportResponse, error) {
	var resp report.MaterialReportResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.ReportResponse, nil
}
