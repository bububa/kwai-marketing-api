package page

import "encoding/json"

// ListRequest 获取魔力建站落地页信息列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主 ID，在获取 access_token 的时候返回
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// ViewComps 包含的组件类型, 多个类型之间是或的关系，0: 图片；1: 文本；2: 表单；3: 按钮；4: 轮播图；5: 视频；6: 地图；7: 应用下载；16: 空白组件；34：小游戏
	ViewComps []int `json:"view_comps,omitempty"`
	// PageName 落地页名称
	PageName string `json:"page_name,omitempty"`
	// PageType 落地页类型，1 表示联盟，非 1 现在都是主站
	PageType int `json:"page_type,omitempty"`
	// FictionIDs 小说 ID 列表，仅对小说行业可选
	FictionIDs []string `json:"fiction_ids,omitempty"`
	// Page 查询的页码数，默认为 1
	Page int `json:"page,omitempty"`
	// PageSize 单页行数，默认为 20，不超过 500
	PageSize int `json:"page_size,omitempty"`
	// ComponentRefIDs 组件中线索通ID（如：小游戏ID），和view_comps类型对应，如：查询含有小游戏id的落地页，view_comps=34;component_ref_ids=123
	ComponentRefIDs []int64 `json:"component_ref_ids,omitempty"`
	// IsPageGroup 是否可创建程序化落地页组，仅对查询可创建程序化的落地页列表有效
	IsPageGroup bool `json:"is_page_group,omitempty"`
	// Select 支持落地页名称模糊查询，落地页ID精准查询，和字段page_name不能同时生效
	Select string `json:"select,omitempty"`
}

// Url implement PostRequest interface
func (r ListRequest) Url() string {
	return "v2/lp/page/list"
}

// Encode implement PostRequest interface
func (r ListRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
