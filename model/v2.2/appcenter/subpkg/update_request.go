package subpkg

import "encoding/json"

type UpdateRequest struct {
	AdvertiserId int64   `json:"advertiser_id,omitempty"`
	PackageId    []int64 `json:"package_id,omitempty"`
	PutStatus    int     `json:"put_status,omitempty"`
}


func (r UpdateRequest) Url() string {
	return "gw/dsp/appcenter/subpkg/mod"
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}