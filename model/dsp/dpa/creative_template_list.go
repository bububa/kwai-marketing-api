package dpa

import "github.com/bububa/kwai-marketing-api/model"

// CreativeTemplateListRequest 获取SDPA创意视频模板 API Request
type CreativeTemplateListRequest struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// OuterID 商品外部ID
	// 填写则返回"canSelect"
	OuterID string `json:"outer_id,omitempty"`
	// AdvertiserID 广告主账号ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// LibraryID 商品库ID
	LibraryID uint64 `json:"library_id,omitempty"`
	// ProductID 商品ID
	// 填写则返回"canSelect"，优先于"outer_id"生效
	ProductID uint64 `json:"product_id,omitempty"`
}

// Url implements PostRequest interface
func (r CreativeTemplateListRequest) Url() string {
	return "gw/dsp/v1/dpa/creative/template/list"
}

// Encode implements PostRequest interface
func (r CreativeTemplateListRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// CreativeTemplateListResponse 获取SDPA创意视频模板 API Response
type CreativeTemplateListResponse struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// TemplateList 创意模板列表
	TemplateList []CreativeTemplate `json:"template_list,omitempty"`
}

// CreativeTemplate 创意模板
type CreativeTemplate struct {
	// Image 封面图URL
	Image string `json:"image,omitempty"`
	// DemoURL 样例视频URL
	DemoURL string `json:"demo_url,omitempty"`
	// TemplateID 模板ID
	TemplateID uint64 `json:"template_id,omitempty"`
	// PackageID 配置包ID
	PackageID uint64 `json:"package_id,omitempty"`
	// TemplateType 模板类型
	// 1-图片模板, 2-视频模板, 3-图片视频模板
	TemplateType int `json:"template_type,omitempty"`
	// CanSelect 是否可选
	CanSelect bool `json:"can_select,omitempty"`
}
