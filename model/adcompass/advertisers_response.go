package adcompass

import "github.com/bububa/kwai-marketing-api/model"

// AdvertisersResponse 罗盘账户
type AdvertisersResponse struct {
	Details []Advertiser `json:"details,omitempty"`
}
type Advertiser struct {
	CorporationName string       `json:"corporation_name,omitempty"`
	AdvertiserName  string       `json:"advertiser_name,omitempty"`
	AdvertiserId    model.Uint64 `json:"advertiser_id,omitempty"`
}
