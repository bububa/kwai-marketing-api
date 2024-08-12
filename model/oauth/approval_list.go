package oauth

import "encoding/json"

// ApprovalListRequest 拉取token下授权广告账户接口 API Request
type ApprovalListRequest struct {
	// AppID 申请应用后快手返回的 app_id
	AppID uint64 `json:"app_id,omitempty"`
	// Secret 申请应用后快手返回的 secret
	Secret string `json:"secret,omitempty"`
	// AccessToken 查询的 access_token
	AccessToken string `json:"access_token,omitempty"`
	// PageNo 分页页码，必填
	PageNo int `json:"page_no,omitempty"`
	// PageSize 分页每页展示条数，必填，最大值为1000
	PageSize int `json:"page_size,omitempty"`
}

// Url implement PostRequest interface
func (r ApprovalListRequest) Url() string {
	return "oauth2/authorize/approval/list"
}

// Encode implement PostRequest interface
func (r ApprovalListRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}

// ApprovalListResponse 拉取token下授权广告账户接口 API Response
type ApprovalListResponse struct {
	Message string `json:"message,omitempty"` // 返回信息
	Data    struct {
		// Details 查询获得的广告主 ID
		Details []uint64 `json:"details,omitempty"`
	} `json:"data,omitempty"` // JSON返回值
	Code int `json:"code,omitempty"` // 返回码
}

// IsError detect if the response is an error
func (r *ApprovalListResponse) IsError() bool {
	return r.Code != 1
}

// Error implement error interface
func (r *ApprovalListResponse) Error() string {
	return r.Message
}
