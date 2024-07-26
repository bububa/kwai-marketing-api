package app

import "github.com/bububa/kwai-marketing-api/model"

// ReleaseListRequest 获取新版应用发布列表【单元创编】API Request
type ReleaseListRequest struct {
	// advertiserID	Long		必填	广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ListType	Integer		可选	列表类型	不传-全部, 1-我创建的, 2-共享给我的
	ListType int `json:"list_type,omitempty"`
	// Platform	String	"ios"	可选	android或ios
	Platform string `json:"platform,omitempty"`
	// AppIDs	Long[]		可选	批量应用id查询	最多支持查询100个
	AppIDs []uint64 `json:"app_ids,omitempty"`
	// KeyWord	String		可选	关键词	支持应用ID或应用名称搜索
	KeyWord string `json:"key_word,omitempty"`
	// Page	Integer		可选	当前页	页码，默认1
	Page int `json:"page,omitempty"`
	// PageSize	Integer		可选	分页大小	个数，默认10
	PageSize int `json:"page_size,omitempty"`
}

func (r ReleaseListRequest) Url() string {
	return "gw/dsp/appcenter/app/release/list"
}

// Encode implement PostRequest interface
func (r ReleaseListRequest) Encode() []byte {
	return model.JSONMarshal(r)
}
