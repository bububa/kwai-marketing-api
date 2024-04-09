package file

import (
	"github.com/bububa/kwai-marketing-api/model"
)

// AdImageListRequest 查询图片接口list接口API Request
type AdImageListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StartDate 开始时间; 与end_date同时传或同时不传；
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间
	EndDate string `json:"end_date,omitempty"`
	// PicTypes 图片类型
	// 0-默认，5-竖版，6-横版，12-开屏。不填，则获取所有类型图片
	PicTypes []int `json:"pic_types,omitempty"`
	// ImageToken 图片 token
	ImageToken string `json:"image_token,omitempty"`
	// ImageTokens 图片 tokens
	ImageTokens []string `json:"image_tokens,omitempty"`
	// Page 请求的页码数
	Page int `json:"page,omitempty"`
	// PageSize 每页行数
	PageSize int `json:"page_size,omitempty"`
}

// Url implement GetRequest interface
func (r AdImageListRequest) Url() string {
	return "v1/file/ad/image/list"
}

// Encode implement GetRequest interface
func (r AdImageListRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// AdImageListResponse 查询图片接口list接口 API Response
type AdImageListResponse struct {
	// TotalCount 图片总数
	TotalCount int `json:"total_count,omitempty"`
	// Details 详情
	Details []Image `json:"details,omitempty"`
}
