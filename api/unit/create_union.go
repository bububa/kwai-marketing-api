package unit

import (
	"context"

	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/unit"
)

// CreateUnion 创建联盟定投广告组
// 【注】白名单账户支持定投联盟广告，创建广告组接口不变，有以下几点限制：
// 1. 仅支持计划类型为“提升应用安装”；“收集销售线索”；“获取电商下单”“提高应用活跃”的广告；
// 2. 仅支持部分定向；3. 转化目标参数convert_id、深度转化目标deep_conversion_type有特殊要求。
func CreateUnion(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.CreateUnionRequest) (uint64, error) {
	var resp unit.CreateUnionResponse
	err := clt.Post(ctx, accessToken, req, &resp)
	if err != nil {
		return 0, err
	}
	return resp.UnitID, nil
}
