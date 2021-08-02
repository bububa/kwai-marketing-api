package target

import "encoding/json"

// TemplateCreateRequest 创建定向模板 API Request
type TemplateCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// TemplateName 定向模板名称; 名字需要唯一
	TemplateName string `json:"template_name,omitempty"`
	// AutoTarget 过滤定向模版; 默认：返回所有类型模版；false：只获取自定义定向模版；true ：只获取智能定向模版
	AutoTarget bool `json:"auto_target,omitempty"`
	// Target 定向信息
	Target *Target `json:"target,omitempty"`
	// BehaviorInterest 行为兴趣定向
	BehaviorInterest *BehaviorInterest `json:"behavior_interest,omitempty"`
}

// Url implement PostRequest interface
func (r TemplateCreateRequest) Url() string {
	return "v1/target/template/create"
}

// Encode implement PostRequest interface
func (r TemplateCreateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
