package comment

import "encoding/json"

// ShieldInfoCreateRequest 增加屏蔽评论信息 API Request
type ShieldInfoCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ShieldContentList 根据 shield_type 类型填写屏蔽信息内容
	ShieldContentList []string `json:"shield_content_list,omitempty"`
	// ShieldType 屏蔽类型
	// 1：评论内容关键词（数量上限 100，长度上限 20），5：用户昵称关键词（数量上限 100，长度上限 20），6：快手 ID（数量上限 200）
	ShieldType int `json:"shield_type,omitempty"`
}

// Url implement PostRequest interface
func (r ShieldInfoCreateRequest) Url() string {
	return "v1/comment/shield_info/create"
}

// Encode implement PostRequest interface
func (r ShieldInfoCreateRequest) Encode() []byte {
	bs, _ := json.Marshal(r)
	return bs
}

// ShieldInfoCreateResponse 增加屏蔽评论信息 API Response
type ShieldInfoCreateResponse struct {
	// ShieldInfoIDList 生成的屏蔽信息id
	// 仅返回创建成功的屏蔽信息id（超过数量限制返回为空；超过长度限制、屏蔽信息重复的屏蔽信息会被过滤掉）
	ShieldInfoIDList []uint64 `json:"shield_info_id_list,omitempty"`
}
