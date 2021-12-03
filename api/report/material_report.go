package report

import (
	"github.com/Shinku-Chen/kwai-marketing-api/core"
	"github.com/Shinku-Chen/kwai-marketing-api/model/report"
)

// MaterialReport 素材报表
func MaterialReport(clt *core.SDKClient, accessToken string, req *report.MaterialReportRequest) (*report.ReportResponse, error) {
	var resp report.ReportResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
