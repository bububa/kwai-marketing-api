package lp

import "encoding/json"

// ConsultListRequest 获取可用咨询组件列表 API Request
type ConsultListRequest struct {
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 每页个数，默认 20
	PageSize int `json:"page_size,omitempty"`
}

// Url implement PostRequest interface
func (r ConsultListRequest) Url() string {
	return "v2/lp/consult/list"
}

// Encode implement PostRequest interface
func (r ConsultListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
