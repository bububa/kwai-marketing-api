package tool

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/tool"
)

// OperationRecordList 账户操作记录信息查询
func OperationRecordList(clt *core.SDKClient, accessToken string, req *tool.OperationRecordListRequest) (*tool.OperationRecordListResponse, error) {
	var resp tool.OperationRecordListResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
