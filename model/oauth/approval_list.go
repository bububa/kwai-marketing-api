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
	// Details 查询获得的广告主 ID
	Details []uint64 `json:"details,omitempty"`
	// 本次查询对应分页是否大于总页数，true-当前分页大于总分页，无需继续向后查询；false-当前分页小于等于总页数，需要查询。
	IsEnd bool `json:"isEnd,omitempty"`
}
