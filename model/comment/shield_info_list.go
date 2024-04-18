package comment

import "encoding/json"

// ShieldInfoListRequest 获取屏蔽评论信息列表 API Request
type ShieldInfoListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 每页条数
	PageSize int `json:"page_size,omitempty"`
	// ShieldTypeList 屏蔽信息类型
	// 1：评论内容关键词，5：用户昵称关键词，6：快手 ID
	ShieldTypeList []int `json:"shield_type_list,omitempty"`
}

// Url implement PostRequest interface
func (r ShieldInfoListRequest) Url() string {
	return "v1/comment/shield_info/list"
}

// Encode implement PostRequest interface
func (r ShieldInfoListRequest) Encode() []byte {
	bs, _ := json.Marshal(r)
	return bs
}

// ShieldInfoListResponse 获取屏蔽评论信息列表 API Response
type ShieldInfoListResponse struct {
	// TotalCount 总条数
	TotalCount int `json:"total_count,omitempty"`
	// Details 屏蔽信息列表
	Details []ShieldInfo `json:"details,omitempty"`
}

type ShieldInfo struct {
	// ShieldInfoID 屏蔽信息ID
	ShieldInfoID uint64 `json:"shield_info_id,omitempty"`
	// ShieldType 屏蔽信息类型
	ShieldType int `json:"shield_type,omitempty"`
	// ShieldContent 屏蔽内容
	ShieldContent string `json:"shield_content,omitempty"`
}
