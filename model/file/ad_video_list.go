package file

import (
	"github.com/bububa/kwai-marketing-api/model"
)

// AdVideoListRequest 查询视频接口list接口API Request
type AdVideoListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PhotoIDs 视频 id列表，不超过 100 个 id
	PhotoIDs []string `json:"photo_ids,omitempty"`
	// NewStatus 视频状态s; 0：删除， -1：全部数据，包含删除 不传默认返回不含删除的数据
	NewStatus *int `json:"new_status,omitempty"`
	// StartDate 开始时间; 与end_date同时传或同时不传；
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间
	EndDate string `json:"end_date,omitempty"`
	// Page 请求的页码数
	Page int `json:"page,omitempty"`
	// PageSize 每页行数
	PageSize int `json:"page_size,omitempty"`
	// OuterLoopNative 是否查询原生视频：0=查询非原生视频；1=查询原生视频。
	OuterLoopNative *int `json:"outer_loop_native,omitempty"`
	// PhotoUserID 视频所属的userId，仅针对原生视频生效，且查询原生视频时必须传递，且仅当视频来源为“本地上传”与“快手客户端（个人主页）”生效。
	PhotoUserID uint64 `json:"photo_user_id,omitempty"`
}

// Url implement PostRequest interface
func (r AdVideoListRequest) Url() string {
	return "v1/file/ad/video/list"
}

// Encode implement PostRequest interface
func (r AdVideoListRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// AdVideoListResponse 查询视频接口list接口 API Response
type AdVideoListResponse struct {
	// TotalCount 图片总数
	TotalCount int `json:"total_count,omitempty"`
	// Details 详情
	Details []Video `json:"details,omitempty"`
}
