package live

import "encoding/json"

type ListJingleBellRequest struct {
	AdvertiserID uint64 `json:"advertiser_id"`
	CurrentPage  int    `json:"current_page"`
	PageSize     int    `json:"page_size"`
	LiveUserID   int64  `json:"live_user_id"`
}

// Url implement PostRequest interface
func (r ListJingleBellRequest) Url() string {
	return "gw/dsp/v1/jingle_bell/live_user_bind_list"
}

// Encode implement PostRequest interface
func (r ListJingleBellRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

type JingleBell struct {
	ConversionType   int    `json:"conversion_type"`
	Active           bool   `json:"active"`
	JingleBellStatus int    `json:"jingle_bell_status"`
	JingleBellId     string `json:"jingle_bell_id"`
	IOSAppID         int64  `json:"ios_app_id"`
	AndroidAppID     int64  `json:"android_app_id"`
}
type ListJingleBellResponse struct {
	TotalCount     int          `json:"total_count"`
	JingleBellList []JingleBell `json:"jingle_bell_list"`
}
