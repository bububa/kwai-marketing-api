package page

import "encoding/json"

// BatchGiveRequest 批量转赠 API Request
type BatchGiveRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// GiveAccountIDs 赠予的目标账户列表
	GiveAccountIDs []uint64 `json:"give_account_ids,omitempty"`
	// PageIDs 赠予的页面list
	PageIDs []uint64 `json:"page_ids,omitempty"`
}

// Url implement PostRequest interface
func (r BatchGiveRequest) Url() string {
	return "gw/magicsite/v1/page/batchGive"
}

// Encode implement PostRequest interface
func (r BatchGiveRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

// BatchGiveResponse 批量转赠 API Response
type BatchGiveResponse struct {
	// Code 1-成功
	Code int `json:"code,omitempty"`
	// ErrorMsg 错误信息
	ErrorMsg string `json:"error_msg,omitempty"`
}

// IsError returns true if response is an error
func (r BatchGiveResponse) IsError() bool {
	return r.Code != 1
}

// Error implements errors.Error interface
func (r BatchGiveResponse) Error() string {
	return r.ErrorMsg
}
