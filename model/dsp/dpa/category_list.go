package dpa

import "github.com/bububa/kwai-marketing-api/model"

// CategoryListRequest 获取商品库类目树 API Request
type CategoryListRequest struct {
	// AdvertiserID 广告主账号ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// LibraryID 商品库ID
	LibraryID uint64 `json:"library_id,omitempty"`
}

// Url implements PostRequest interface
func (r CategoryListRequest) Url() string {
	return "gw/dsp/v1/dpa/category/list"
}

// Encode implements PostRequest interface
func (r CategoryListRequest) Encode() []byte {
	return model.JSONMarshal(r)
}

// CategoryListResponse 获取商品库类目树 API Response
type CategoryListResponse struct {
	// Details 类目信息列表
	Details []Category `json:"details,omitempty"`
}

// Category 类目信息
type Category struct {
	// Label 类目名称
	Label string `json:"label,omitempty"`
	// Children 子类目
	Children []Category `json:"children,omitempty"`
	// Value 类目ID
	Value uint64 `json:"value,omitempty"`
}
