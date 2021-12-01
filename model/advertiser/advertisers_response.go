package advertiser

// AdvertisersResponse 罗盘账户
type AdvertisersResponse struct {
	
	Details []Advertiser `json:"details,omitempty"`
}
type Advertiser struct {
	CorporationName string `json:"corporation_name,omitempty"`
	AdvertiserName  string `json:"advertiser_name,omitempty"`
	AdvertiserId    int64    `json:"advertiser_id,omitempty"`
}
