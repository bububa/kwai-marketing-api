package subpkg

import "encoding/json"

type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	AppId        int64  `json:"app_id"`
	KeyWord      string `json:"key_word"`
	Page         int    `json:"page"`
	PageSize     int    `json:"page_size"`
}

func (r ListRequest) Url() string {
	return "gw/dsp/appcenter/subPackage/release/list"
}

// Encode implement PostRequest interface
func (r ListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
