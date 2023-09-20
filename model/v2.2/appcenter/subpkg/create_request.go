package subpkg

import "encoding/json"

// CreateRequest 创建分包 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	//parent_package_id	Long		必填	应用（母）包id	仅支持android应用的包id新建分包
	ParentPackageID uint64 `json:"parent_package_id,omitempty"`
	//type	Integer		必填	分包方式	1-系统自动分包，2-上传渠道号列表
	Type int `json:"type,omitempty"`
	//count	Integer		可选	分包数量	当type=1时填写，单次最多100
	Count int `json:"count,omitempty"`
	//channel_id	String[]		可选	上传的渠道号列表	当type=2时填写，单次最多填写100个。同一应用包下填写的渠道号不可重复
	ChannelID []string `json:"channel_id,omitempty"`
}

func (r CreateRequest) Url() string {
	return "gw/dsp/appcenter/subpkg/add"
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
