package creative

import (
	"net/url"
	"strconv"
)

// CreativeTagAdviseRequest 创意标签填写建议 API Request
type CreativeTagAdviseRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Url implement GetRequest interface
func (r CreativeTagAdviseRequest) Url() string {
	return "v1/creative/creative_tag/advise"
}

// Encode implement GetRequest interface
func (r CreativeTagAdviseRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	return values.Encode()
}
