package live

import "encoding/json"

type ListLiveUserRequest struct {
	AdvertiserID     uint64 `json:"advertiser_id"`
	CurrentPage      int    `json:"current_page"`
	PageSize         int    `json:"page_size"`
	AuditStatusList  []int  `json:"audit_status_list"`
	LiveUserTypeList []int  `json:"live_user_type_list"`
	LiveUserIds      []int  `json:"live_user_ids"`
}

// Url implement PostRequest interface
func (r ListLiveUserRequest) Url() string {
	return "gw/dsp/v1/live_manage/get_live_users"
}

// Encode implement PostRequest interface
func (r ListLiveUserRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

type ListLiveUserResponse struct {
	TotalCount   int        `json:"total_count"`
	LiveUserList []LiveUser `json:"live_user_list"`
}

type LiveUser struct {
	LiveUserType int   `json:"live_user_type"`
	Living       bool  `json:"living"`
	AccountId    int64 `json:"account_id"`
	AuditStatus  int   `json:"audit_status"`
	LiveUserId   int64 `json:"live_user_id"`
}
