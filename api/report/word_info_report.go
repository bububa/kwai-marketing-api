package report

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/report"
)

// WordInfoReport
func WordInfoReport(clt *core.SDKClient, accessToken string, req *report.WordInfoReportRequest) (*report.WordInfoReport, error) {
	var resp report.WordInfoReportResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
