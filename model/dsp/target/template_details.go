package target

import (
	"github.com/bububa/kwai-marketing-api/model"
	"github.com/bububa/kwai-marketing-api/model/dsp/unit"
)

// TemplateDetailsRequest 查询定向模板 API Request
type TemplateDetailsRequest struct {
	// TemplateID 模板ID列表（默认ID和分页参数二选一必传）
	TemplateID []uint64 `json:"template_id,omitempty"`
	// AdvertiserID 广告主 ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 请求的页码，默认为 1（page和page_size要么都传，要么都不传；都不传时需要传模板ID）
	Page int `json:"page,omitempty"`
	// PageSize 每页行数，默认 20
	PageSize int `json:"page_size,omitempty"`
}

// Url implement GetRequest interface
func (r TemplateDetailsRequest) Url() string {
	return "gw/dsp/target/template/details"
}

// Encode implement GetRequest interface
func (r TemplateDetailsRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// TemplateDetailsResponse 查询定向模板 API Response
type TemplateDetailsResponse struct {
	Details []Template `json:"details,omitempty"`
	// TotalCount 总数
	TotalCount int `json:"total_count,omitempty"`
	// CurrentPage 当前页码, 第一页是1
	CurrentPage int `json:"current_page"`
	// PageSize 分页大小
	PageSize int `json:"page_size"`
}

// Template 定向模板
type Template struct {
	// Target 定向信息
	Target *unit.Target `json:"target,omitempty"`
	// TemplateName 定向模板名称
	TemplateName string `json:"template_name,omitempty"`
	// CreateTime 定向模板创建时间，"2019-06-11 15:17:25"
	CreateTime string `json:"create_time,omitempty"`
	// UpdateTime 定向模板最近修改时间
	UpdateTime string `json:"update_time,omitempty"`
	// TemplateID 定向模板 ID
	TemplateID uint64 `json:"template_id,omitempty"`
	// UnitCount 绑定unit个数
	UnitCount int `json:"unit_count,omitempty"`
}
