package adcompass

import (
	"encoding/json"
)

// QuotaTendingRequest 磁力罗盘对外 quota 腾挪接口
type QuotaTendingRequest struct {
	// TendingDetails 待腾挪组
	TendingDetails []TendingDetails `json:"tending_details,omitempty"`
}

// TendingDetails 待腾挪组
type TendingDetails struct {
	// FromAccountID 转出 account_id
	FromAccountID uint64 `json:"from_account_id,omitempty"`
	// ToAccountID 转入 account_id
	ToAccountID uint64 `json:"to_account_id,omitempty"`
	// TendingCount 腾挪额度
	TendingCount int64 `json:"tending_count,omitempty"`
}

// Url implement PostRequest interface
func (r QuotaTendingRequest) Url() string {
	return "gw/ad_compass/quota/tending"
}

// Encode implement PostRequest interface
func (r QuotaTendingRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
