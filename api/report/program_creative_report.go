package report

import (
	"github.com/Shinku-Chen/kwai-marketing-api/core"
	"github.com/Shinku-Chen/kwai-marketing-api/model/report"
)

// ProgramCreativeReport 程序化创意2.0数据
func ProgramCreativeReport(clt *core.SDKClient, accessToken string, req *report.ProgramCreativeReportRequest) (*report.ReportResponse, error) {
	var resp report.ReportResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
