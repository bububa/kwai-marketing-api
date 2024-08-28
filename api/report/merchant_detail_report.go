package report

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/report"
)

// MerchantDetailReport 小店通转化数据报表
func MerchantDetailReport(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.MerchantDetailReportRequest) (*report.MerchantDetailReportResponse, error) {
	var resp report.MerchantDetailReportResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
