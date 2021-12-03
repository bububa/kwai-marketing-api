package report

import (
	"github.com/Shinku-Chen/kwai-marketing-api/core"
	"github.com/Shinku-Chen/kwai-marketing-api/model/report"
)

// AgentReport 代理商数据（t-1数据需要第二天9点以后获取）
func AgentReport(clt *core.SDKClient, accessToken string, req *report.AgentReportRequest) (*report.AgentReportResponse, error) {
	var resp report.AgentReportResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
