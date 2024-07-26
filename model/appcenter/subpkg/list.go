package subpkg

import "github.com/bububa/kwai-marketing-api/model"

// ListRequest 【应用中心】获取分包管理/回收站列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AppID 应用id
	AppID uint64 `json:"app_id,omitempty"`
	// ListType 列表类型
	// 不填为分包管理列表，填2-分包回收列表
	ListType int `json:"list_type,omitempty"`
	// KeyWord 关键词
	// 支持渠道号关键词搜索
	KeyWord string `json:"key_word,omitempty"`
	// Status 	分包状态筛选
	// 分包管理列表生效，筛选分包状态，单选。不传默认全部，0-全部，1-审核中，2-审核未通过，4-已发布，6-分包构建中，7-分包更新中，8-分包构建失败
	Status int `json:"status,omitempty"`
	// Version 分包版本筛选
	// 分包管理列表生效，筛选版本信息，多选。不传默认全部。
	Version []string `json:"version,omitempty"`
	// Page 当前页	页码，默认1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小	个数，默认10
	PageSize int `json:"page_size,omitempty"`
}

// Url implements PostRequest interface
func (r ListRequest) Url() string {
	return "gw/dsp/appcenter/subPackage/list"
}

// Encode implement PostRequest interface
func (r ListRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// ListResponse 获取新版分包列表 API Response
type ListResponse struct {
	// CurrentPage 当前页
	CurrentPage int `json:"current_page,omitempty"`
	// PageSize 分页大小
	PageSize int `json:"page_size,omitempty"`
	// TotalCount 总数
	TotalCount int `json:"total_count,omitempty"`
	// List 应用分包列表
	List []SubPackage `json:"list,omitempty"`
}
