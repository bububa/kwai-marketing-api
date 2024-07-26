package subpkg

import "github.com/bububa/kwai-marketing-api/model"

// AddRequest 创建分包 API Request
type AddRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ParentPackageID	Long		必填	应用（母）包id	仅支持android应用的包id新建分包
	ParentPackageID uint64 `json:"parent_package_id,omitempty"`
	// Type	Integer		必填	分包方式	1-系统自动分包，2-上传渠道号列表
	Type int `json:"type,omitempty"`
	// Count	Integer		可选	分包数量	当type=1时填写，单次最多100
	Count int `json:"count,omitempty"`
	// ChannelID	String[]		可选	上传的渠道号列表	当type=2时填写，单次最多填写100个。同一应用包下填写的渠道号不可重复
	ChannelID []string `json:"channel_id,omitempty"`
}

func (r AddRequest) Url() string {
	return "gw/dsp/appcenter/subpkg/add"
}

// Encode implement PostRequest interface
func (r AddRequest) Encode() []byte {
	return model.JSONMarshal(r)
}
