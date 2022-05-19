package tool

import "encoding/json"

// OperationRecordListRequest 账户操作记录信息查询 API Request
type OperationRecordListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// OperationType 操作类型; 1：新增2：修改删除（可通过content_log： "update_data": "删除"查看）
	OperationType int `json:"operation_type,omitempty"`
	// OperationTarget 操作目标类型; 目前只支持：1:账户2：计划3：广告组4：创意5：视频6：app应用7：人群包
	OperationTarget int `json:"operation_target,omitempty"`
	// RoleType 操作人; 1：广告主2：代理商3：系统4：管理员5：Market Api
	RoleType int `json:"role_type,omitempty"`
	// Page 页码数; 默认为1
	Page int `json:"page,omitempty"`
	// PageSize 单页行数; 默认为20，不超过500
	PageSize int `json:"page_size,omitempty"`
	// StartDate 开始时间; 最多查询最近6个月的操作记录
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间
	EndDate string `json:"end_date,omitempty"`
}

// Url implement PostRequest interface
func (r OperationRecordListRequest) Url() string {
	return "v1/tool/operation_record/list"
}

// Encode implement PostRequest interface
func (r OperationRecordListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
