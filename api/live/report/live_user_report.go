package report

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/live/report"
)

// ListLiveUserReport 直播间报表
func ListLiveUserReport(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.ListLiveUserReportRequest) (*report.ListLiveReportResponse, error) {
	var resp report.ListLiveReportResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
