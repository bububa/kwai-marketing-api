package dpa

import "github.com/bububa/kwai-marketing-api/model"

// TemplateListRequest 查询 DPA 模板信息 API Request
type TemplateListRequest struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// TemplateTypes 模板类型
	// 模板类型:1:竖版图片、2:横版图片3:横版视频、4:竖版视频
	TemplateTypes []int `json:"template_types,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Url implements PostRequest interface
func (r TemplateListRequest) Url() string {
	return "gw/dsp/dpa/template/list"
}

// Encode implement PostRequest interface
func (r TemplateListRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// TemplateListResponse 查询 DPA 模板信息 API Response
type TemplateListResponse struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// TemplateList 模板列表
	TemplateList []*Template `json:"template_list,omitempty"`
}
