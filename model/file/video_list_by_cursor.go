package file

import (
	"github.com/bububa/kwai-marketing-api/model"
)

// VideoListByCursorRequest 查询视频接口list接口API Request
// 本接口适用于全量查询账号内的视频，通过游标的方式不断获取数据，性能和稳定性都比list分页接口高。 不支持条件筛选。
type VideoListByCursorRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Cursor 游标
	Cursor *int64 `json:"cursor,omitempty"`
	// Limit 请求的页码数
	Limit int `json:"limit,omitempty"`
}

// Url implement PostRequest interface
func (r VideoListByCursorRequest) Url() string {
	return "gw/dsp/video/listByCursor"
}

// Encode implement PostRequest interface
func (r VideoListByCursorRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// VideoListByCursorResponse 查询视频接口list接口 API Response
type VideoListByCursorResponse struct {
	// TotalCount 图片总数
	TotalCount int `json:"total_count,omitempty"`
	// Details 详情
	Details []Video `json:"details,omitempty"`
}
