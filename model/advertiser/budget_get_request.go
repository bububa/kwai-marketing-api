package advertiser

import (
	"net/url"
	"strconv"
)

// BudgetGetRequest 账户日预算查询APIRequest
type BudgetGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Url implement GetRequest interface
func (r BudgetGetRequest) Url() string {
	return "v1/advertiser/budget/get"
}

// Encode implement GetRequest interface
func (r BudgetGetRequest) Encode() string {
	values := url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	return values.Encode()
}
