package app

import "github.com/bububa/kwai-marketing-api/model"

// ListRequest 【应用中心】获取应用列表 API Request
type ListRequest struct {
	// AdvertiserID	Long		必填	广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ListType	Integer		可选	列表类型	不传-全部, 1-我创建的, 2-共享给我的
	ListType int `json:"list_type,omitempty"`
	// Platform	String	"ios"	可选	android或ios
	Platform string `json:"platform,omitempty"`
	// AppStatus 应用状态
	// 不传-全部，1-审核中，2-审核未通过，3-待发布，4-已发布，5-已下架
	AppStatus int `json:"app_status,omitempty"`
	// KeyWord	String		可选	关键词	支持应用ID或应用名称搜索
	KeyWord string `json:"key_word,omitempty"`
	// StartDate 发布时间范围-起始
	// 同时需要填写end_date
	StartDate string `json:"start_date,omitempty"`
	// EndDate 发布时间范围-截止
	// 同时需要填写start_date
	EndDate string `json:"end_date,omitempty"`
	// Page	Integer		可选	当前页	页码，默认1
	Page int `json:"page,omitempty"`
	// PageSize	Integer		可选	分页大小	个数，默认10
	PageSize int `json:"page_size,omitempty"`
}

// Url implements PostRequest interface
func (r ListRequest) Url() string {
	return "gw/dsp/appcenter/app/list"
}

// Encode implement PostRequest interface
func (r ListRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// ListRequest 【应用中心】获取应用列表 API Response
type ListResponse struct {
	// CurrentPage 当前页
	CurrentPage int `json:"current_page,omitempty"`
	// PageSize 分页大小
	PageSize int `json:"page_size,omitempty"`
	// TotalCount 总数
	TotalCount int `json:"total_count,omitempty"`
	// List 	应用列表
	List []App `json:"list,omitempty"`
}
