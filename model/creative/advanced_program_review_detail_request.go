package creative

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// AdvancedProgramReviewDetailRequest 获取程序化创意2.0审核信息
type AdvancedProgramReviewDetailRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UnitIDs 广告组ID;数量小于等于20个
	UnitIDs []uint64 `json:"unit_ids,omitempty"`
}

// Url implement GetRequest interface
func (r AdvancedProgramReviewDetailRequest) Url() string {
	return "v2/creative/advanced/program/review_detail"
}

// Encode implement GetRequest interface
func (r AdvancedProgramReviewDetailRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	unitIds, _ := json.Marshal(r.UnitIDs)
	values.Set("unit_ids", string(unitIds))
	return values.Encode()
}
