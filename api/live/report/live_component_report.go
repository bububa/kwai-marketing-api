package report

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/live/report"
)

// ListLiveComponentReport 直播间组件报表
func ListLiveComponentReport(clt *core.SDKClient, accessToken string, req *report.ListLiveComponentReportRequest) (*report.ListLiveComponentReportResponse, error) {
	var resp report.ListLiveComponentReportResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
