package subpkg

import "github.com/bububa/kwai-marketing-api/model"

// ReleaseListRequest 获取新版分包发布列表【单元创编】 API Request
type ReleaseListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AppID 应用id
	AppID uint64 `json:"app_id,omitempty"`
	// KeyWord 关键词
	// 支持渠道号关键词搜索
	KeyWord string `json:"key_word,omitempty"`
	// Page 当前页	页码，默认1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小	个数，默认10
	PageSize int `json:"page_size,omitempty"`
}

func (r ReleaseListRequest) Url() string {
	return "gw/dsp/appcenter/subPackage/release/list"
}

// Encode implement PostRequest interface
func (r ReleaseListRequest) Encode() []byte {
	return model.JSONMarshal(r)
}
