package advertiser

import "encoding/json"

// BudgetGetRequest 账户日预算查询APIRequest
type BudgetGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Url implement GetRequest interface
func (r BudgetGetRequest) Url() string {
	return "v1/advertiser/budget/get"
}

// Encode implement PostRequest interface
func (r BudgetGetRequest) Encode() []byte {
	bs, _ := json.Marshal(r)
	return bs
}
