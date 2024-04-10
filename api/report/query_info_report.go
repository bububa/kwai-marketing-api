package report

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/report"
)

// QueryInfoReport 搜索词报表
func QueryInfoReport(clt *core.SDKClient, accessToken string, req *report.QueryInfoReportRequest) (*report.QueryInfoReport, error) {
	var resp report.QueryInfoReportResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
