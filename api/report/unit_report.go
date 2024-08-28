package report

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/report"
)

// UnitReport 广告组数据
func UnitReport(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.UnitReportRequest) (*report.ReportResponse, error) {
	var resp report.ReportResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
