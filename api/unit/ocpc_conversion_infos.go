package unit

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/unit"
)

// OcpcConversionInfos 获取可选的深度转化目标
// 新门槛标准： 1.应用下载类—— 1)账户维度，相应的 event_type 回传 >=1; 或2)应用包维度，用转化追踪联调通过相应的 event_type 2.落地页类 账户维度，相应的 event_type 回传 >=1;
func OcpcConversionInfos(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.OcpcConversionInfosRequest) (*unit.OcpcConversionInfosResponse, error) {
	var resp unit.OcpcConversionInfosResponse
	err := clt.Get(ctx, accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
