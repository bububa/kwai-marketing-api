package advertiser

// AdvertisersResponse 罗盘账户
type AdvertisersResponse struct {
	// Balance 账户总余额;单位：元
	Details []Advertiser `json:"details,omitempty"`
}
type Advertiser struct {
	CorporationName string `json:"corporation_name,omitempty"`
	AdvertiserName  string `json:"advertiser_name,omitempty"`
	AdvertiserId    int    `json:"advertiser_id,omitempty"`
}
