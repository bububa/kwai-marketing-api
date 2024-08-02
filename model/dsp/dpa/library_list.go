package dpa

import "github.com/bububa/kwai-marketing-api/model"

// LibraryListRequest 获取商品库列表 API Request
type LibraryListRequest struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// Name 商品库名称
	Name string `json:"name,omitempty"`
	// BizID 商品库业务类型 1-主站, 2-联盟, 3-通用
	BizID []int `json:"biz_id,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// LibraryID 商品库ID
	LibraryID uint64 `json:"library_id,omitempty"`
	// Status 商品库状态 1-审核中, 2-使用中, 3-审核失败, 4-暂停使用, 5-XML初始化中, 6-XML初始化失败
	Status int `json:"status,omitempty"`
	// QueryType 商品库权限类型 1-使用权限, 2-编辑权限(含使用权限)
	QueryType int `json:"query_type,omitempty"`
}

// Url implement PostRequest interface
func (r LibraryListRequest) Url() string {
	return "gw/dsp/dpa/library/list"
}

// Encode implement PostRequest interface
func (r LibraryListRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// LibraryListResponse 获取商品库列表 API Response
type LibraryListResponse struct {
	// PageInfo 列表页参数
	PageInfo *model.PageInfo         `json:"page_info,omitempty"`
	Data     []AdDpaLibraryViewSneak `json:"data,omitempty"`
}

// AdDpaLibraryViewSneak 商品库信息
type AdDpaLibraryViewSneak struct {
	// Name 商品库名称
	Name string `json:"name,omitempty"`
	// LibraryDesc 商品库描述
	LibraryDesc string `json:"library_desc,omitempty"`
	// LibraryID 商品库ID
	LibraryID uint64 `json:"library_id,omitempty"`
	// Status 商品库状态 1-审核中, 2-使用中, 3-审核失败, 4-暂停使用, 5-XML初始化中, 6-XML初始化失败
	Status int `json:"status,omitempty"`
	// CreateTime 商品库创建时间 毫秒时间戳
	CreateTime int64 `json:"create_time,omitempty"`
	// BizID 商品库业务类型 0-未知, 1-主站, 2-联盟, 3-通用
	BizID int `json:"biz_id,omitempty"`
}
